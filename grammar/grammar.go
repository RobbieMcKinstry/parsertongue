package grammar

import (
	"io"
	"os"
	"unicode"

	"golang.org/x/exp/ebnf"
)

// G is a grammar
type G struct {
	// the underlying grammar
	gram ebnf.Grammar
	// lexemes is the list of Names of the strings
	// representing lexical productions
	lexemes []string

	// productions is the list of Names of the strings
	// representing syntactic productions
	prods []string

	// open is a mockable version of os.Open
	open func(string) (io.Reader, error)
}

func wrapOpen() func(string) (io.Reader, error) {
	return func(path string) (io.Reader, error) {
		f, err := os.Open(path)
		return f, err
	}
}

// Grammar returns the grammar
func (gram *G) Grammar() ebnf.Grammar {
	return gram.gram
}

// LexemeNames returns the lex names
func (gram *G) LexemeNames() []string {
	return gram.lexemes
}

// ProdNames returns the prodnames
func (gram *G) ProdNames() []string {
	return gram.prods
}

// Prod returns the production with the
// given name
func (gram *G) Prod(name string) *ebnf.Production {
	prod, ok := gram.gram[name]
	if !ok {
		return nil
	}
	return prod
}

// IsLexeme returns true if the name
// starts with a lower case letter
func IsLexeme(name string) bool {
	if len(name) < 1 {
		return false
	}

	firstLetter := []rune(name)[0]
	return unicode.IsLower(firstLetter)
}

// Children returns the list of names reprenting the children
// of this tree node
func Children(parent *ebnf.Production) []string {
	return []string{}
}

func newG() *G {
	return &G{
		open: wrapOpen(),
	}
}

func (gram *G) init(path string) {

	file, err := gram.open(path)
	if err != nil {
		panic(err)
	}

	grammar, err := ebnf.Parse(path, file)
	if err != nil {
		panic(err)
	}
	gram.gram = grammar
	gram.lexemes = lowercaseProds(grammar)
	gram.prods = uppercaseProds(grammar)
}

func lowercaseProds(gram ebnf.Grammar) []string {
	if gram == nil {
		return []string{}
	}

	var res []string
	for name := range gram {
		if IsLexeme(name) {
			res = append(res, name)
		}
	}
	return res
}

func uppercaseProds(gram ebnf.Grammar) []string {
	if gram == nil {
		return []string{}
	}

	var res []string
	for name := range gram {
		if !IsLexeme(name) {
			res = append(res, name)
		}
	}
	return res
}
