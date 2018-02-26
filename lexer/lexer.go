package lexer

import (
	"golang.org/x/exp/ebnf"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
	"github.com/RobbieMcKinstry/parsertongue/system/formatter"
	"github.com/sirupsen/logrus"
)

const eof = rune(0)

// L is the lexer
type L struct {
	log    *logrus.Logger
	gram   *grammar.G
	reader CloneScanner
	out    chan<- Token
}

// NewLexer is the lexer constructor
func NewLexer(gram *grammar.G, data []byte, tokenStream chan<- Token) *L {
	scanner := NewBufferScanner(data)
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.Formatter = formatter.Colored{}
	return &L{
		log:    log,
		gram:   gram,
		reader: scanner,
		out:    tokenStream,
	}
}

// Clone this lexer
func (lex *L) Clone() *L {
	next := new(L)
	next.gram = lex.gram
	next.reader = lex.reader.Clone()
	next.out = lex.out
	return next
}

// Lex returns the lexer and the token stream
func Lex(gram *grammar.G, data []byte) (*L, <-chan Token) {
	channel := make(chan Token)
	lexer := NewLexer(gram, data, channel)
	go lexer.run()
	return lexer, channel
}

// run will calculate the lexeme DAG and generate lexemes of those types
func (lex *L) run() {
	// First, first the entrant productions...
	lex.log.Debugf("Making entrant prods")
	entrantNames := grammar.FindEntrantProds(lex.gram)
	// collect the actual productions from the grammar
	lex.log.Debugf("Collecting prods by name")
	prods := lex.collectProdsByName(entrantNames)

	lex.log.Debugf("Making state fns")
	stateFns := lex.makeStateFns(prods)

	// now, combine those state funcs with the state funcs
	// generated for the token literals found in the
	// non-lexical productions
	// TODO not implemented at this time.
	tokenLiterals := lex.gram.FindTokenLiterals()
	for _, tok := range tokenLiterals {
		tokenStateFn := lex.makeToken(tok)
		stateFns = append(stateFns, tokenStateFn)
	}

	lex.log.Debugf("StateFns created. Beginning to lex all prods")

	for {
		var token = Token{}

		prod, count := lex.maxProds(prods, stateFns)
		if count == -1 {
			lex.log.Debugf("No prods match. Breaking…")
			close(lex.out)
			break
		}
		if prod == nil {
			token.IsLexemeLiteral = true
			lex.log.Debugf("Make len is %v of type 'lexeme literal'", count)
		} else {
			token.typ = prod
			lex.log.Debugf("Make len is %v of type %s", count, prod.Name.String)
		}

		lexeme := make([]rune, 0, count)
		// capture the lexeme in a string
		for i := 0; i < count; i++ {
			lexeme = append(lexeme, lex.next())
		}

		token.Val = string(lexeme)
		lex.out <- token

		// Now, exhaust any remaining whitespace.
		lex.clearWhitespace()
		if lex.peek() == eof {
			close(lex.out)
			break
		}
	}
}

func (lex *L) makeStateFns(prods []*ebnf.Production) []StateFn {
	fns := make([]StateFn, 0, len(prods))
	for _, prod := range prods {
		fns = append(fns, lex.toStateFn(prod.Expr))
	}
	return fns
}

// maxProds returns the prod with the longest count.
// FIXME: RETURNS NIL IF THE PRODUCTION IS AN ENTRANT LEXICAL LITERAL
// A more robust solution is to have an interface type that either
// returns a production OR it returns a lexeme literal.
func (lex *L) maxProds(prods []*ebnf.Production, fns []StateFn) (*ebnf.Production, int) {
	counts := make([]int, 0, len(prods))
	for _, fn := range fns {
		count := fn.Exhaust(lex.Clone())
		counts = append(counts, count)
	}

	index := max(counts)
	var production *ebnf.Production
	if index < len(prods) {
		production = prods[index]
	}
	return production, counts[index]
}

func max(slice []int) int {
	maxIndex := 0
	for i, val := range slice {
		if val > slice[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func (lex *L) clearWhitespace() {
	var count = StateFn(Whitespace).Exhaust(lex.Clone())
	lex.advance(count)
}

func (lex *L) collectProdsByName(names []string) []*ebnf.Production {
	var prods []*ebnf.Production
	for _, name := range names {
		prods = append(prods, lex.gram.Prod(name))
	}
	return prods
}

// next will return the next rune, returning eof if there is not next rune
func (lex *L) next() rune {
	r, _, err := lex.reader.ReadRune()
	if err != nil {
		return eof
	}
	return r
}

func (lex *L) peek() rune {
	r := lex.next()
	lex.backup()
	return r
}

func (lex *L) backup() {
	lex.reader.UnreadRune()
}

func (lex *L) advance(pos int) {
	for i := 0; i < pos; i++ {
		lex.next()
	}
}
