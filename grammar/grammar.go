package grammar

import (
	"io"
	"os"
	"unicode"

	"golang.org/x/exp/ebnf"
)

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

func (gram G) Grammar() ebnf.Grammar {
	return gram.gram
}

func (gram G) LexemeNames() []string {
	return gram.lexemes
}

func (gram G) ProdNames() []string {
	return gram.prods
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
		if len(name) < 1 {
			continue
		}

		firstLetter := []rune(name)[0]
		if unicode.IsLower(firstLetter) {
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
		if len(name) < 1 {
			continue
		}

		firstLetter := []rune(name)[0]
		if !unicode.IsLower(firstLetter) {
			res = append(res, name)
		}
	}
	return res
}
