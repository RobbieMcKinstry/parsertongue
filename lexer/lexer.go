package lexer

import (
	"golang.org/x/exp/ebnf"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
)

const eof = rune(0)

// L is the lexer
type L struct {
	gram   *grammar.G
	reader CloneScanner
	out    chan<- Token
}

// NewLexer is the lexer constructor
func NewLexer(gram *grammar.G, data []byte, tokenStream chan<- Token) *L {
	scanner := NewBufferScanner(data)
	return &L{
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

func lex(gram *grammar.G, data []byte) (*L, <-chan Token) {
	channel := make(chan Token)
	lexer := NewLexer(gram, data, channel)
	go lexer.run()
	return lexer, channel
}

// run will calculate the lexeme DAG and generate lexemes of those types
func (lex *L) run() {
	// First, first the entrant productions...
	entrantNames := grammar.FindEntrantProds(lex.gram)
	// collect the actual productions from the grammar
	prods := lex.collectProdsByName(entrantNames)
	_ = prods

	// now, for each of these prods, which which one produces the longest result...
	// TODO
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
