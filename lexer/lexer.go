package lexer

import (
	"io"

	"golang.org/x/exp/ebnf"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
)

// L is the lexer
type L struct {
	gram   *grammar.G
	reader io.Reader
	out    chan<- Token
}

// NewLexer returns a lexer
func NewLexer(gram *G, r io.Reader, tokenStream chan<- Token) *L {
	return &L{
		gram:   gram,
		reader: r,
		out:    tokenStream,
	}
}

func lex(gram *G, r io.Reader) (*L, <-chan Token) {
	channel := make(chan Token)
	lexer := NewLexer(gram, r, channel)
	go lexer.run()
	return lexer, channel
}

// run will calculate the lexeme DAG and generate lexemes of those types
func (lex *L) run() {
	// First, first the entrant productions...
	entrantNames := FindEntrantProds(lex.gram)
	// collect the actual productions from the grammar
	prods := collectProdsByName(entrantnames)

	// now, for each of these prods, which which one produces the longest result...
	// TODO
}

func (lex *L) collectProdsByName(names []string) []*ebnf.Production {
	var prods []*ebnf.Production
	for _, name := range names {
		prods = append(prods, lex.gram[name])
	}
	return prods
}
