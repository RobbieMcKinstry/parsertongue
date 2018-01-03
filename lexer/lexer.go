package lexer

import (
	"fmt"

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

// Lex returns the lexer and the token stream
func Lex(gram *grammar.G, data []byte) (*L, <-chan Token) {
	channel := make(chan Token)
	fmt.Println("Making a new lexer")
	lexer := NewLexer(gram, data, channel)
	fmt.Println("Calling run in its own goroutine")
	go lexer.run()
	return lexer, channel
}

func (lex *L) findEntrantProds() []string {
	entrantNames := grammar.FindEntrantProds(lex.gram)
	printEntrantProds(entrantNames)

	return entrantNames
}

func printEntrantProds(entrantNames []string) {
	fmt.Println("Entrant Prods:")
	for _, prod := range entrantNames {
		fmt.Printf("\t%v\n", prod)
	}
}

func (lex *L) entrantStateFns() []StateFn {
	// First, first the entrant productions...
	entrantNames := lex.findEntrantProds()

	// collect the actual productions from the grammar
	fmt.Println("Collecting prods by name")
	prods := lex.collectProdsByName(entrantNames)

	fmt.Println("Making state fns")
	return lex.makeStateFns(prods)
}

func (lex *L) tokenStateFns() []StateFn {

	var fns []StateFn
	tokenLiterals := lex.gram.FindTokenLiterals()
	printTokens(tokenLiterals)

	for _, tok := range tokenLiterals {
		tokenStateFn := lex.makeToken(tok)
		fns = append(fns, tokenStateFn)
	}
	return fns
}

func printTokens(toks []*ebnf.Token) {
	fmt.Println("Tokens:")
	for _, tok := range toks {
		fmt.Printf("\t%v\n", tok.String)
	}
}

// run will calculate the lexeme DAG and generate lexemes of those types
func (lex *L) run() {
	entrantStateFns := lex.entrantStateFns()
	tokenStateFns := lex.tokenStateFns()

	var stateFns []StateFn
	stateFns = append(stateFns, entrantStateFns...)
	stateFns = append(stateFns, tokenStateFns...)

	entrantNames := lex.findEntrantProds()
	prods := lex.collectProdsByName(entrantNames)
	// now, combine those state funcs with the state funcs
	// generated for the token literals found in the
	// non-lexical productions
	// TODO not implemented at this time.
	fmt.Println("StateFns created. Beginning to lex all prods")

	for {
		lex.clearWhitespace()
		prod, count := lex.maxProds(prods, stateFns)
		var isTokenLiteral = false
		if count == -1 {
			fmt.Println("No prods match. Breaking…")
			close(lex.out)
			break
		} else if prod == nil {
			// Then this is token literal
			isTokenLiteral = true
		}

		if isTokenLiteral {
			fmt.Printf("Make len is %v of type %s\n", count, "token")
		} else {
			fmt.Printf("Make len is %v of type %s\n", count, prod.Name.String)
		}
		lexeme := make([]rune, 0, count)
		// capture the lexeme in a string
		for i := 0; i < count; i++ {
			lexeme = append(lexeme, lex.next())
		}

		tok := Token{
			typ:      prod,
			Val:      string(lexeme),
			IsLexeme: isTokenLiteral,
		}
		fmt.Printf("Found token %v\n", string(lexeme))

		lex.out <- tok

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
func (lex *L) maxProds(prods []*ebnf.Production, fns []StateFn) (*ebnf.Production, int) {
	counts := make([]int, 0, len(fns))
	for _, fn := range fns {
		count := fn.Exhaust(lex.Clone())
		counts = append(counts, count)
	}

	index := max(counts)
	if index >= len(prods) {
		// then this is a lexical production.
		return nil, counts[index]
	}

	return prods[index], counts[index]
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
