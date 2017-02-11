package main

import (
	"log"
	"os"

	"golang.org/x/exp/ebnf"
)

const (
	FileName     = "./grammar.ebnf"
	SentenceRoot = "SourceFile"
)

func OpenGrammar() ebnf.Grammar {

	file, err := os.Open(FileName)
	if err != nil {
		log.Fatal(err)
	}

	grammar, err := ebnf.Parse(FileName, file)
	if err != nil {
		log.Fatal(err)
	}
	return grammar
}

func VerifyGrammar(g ebnf.Grammar) ebnf.Grammar {
	if err := ebnf.Verify(g, SentenceRoot); err != nil {
		log.Fatal(err)
	}
	return g
}

var prebuilt = map[string]stateFn{
	"newline": stateFn(NewlineStateFn),
	//	"unicode_char": ,
	//	"unicode_letter": ,
	//	"unicode_digit": ,
	//	"token": ,
}

func lookupPrebuilt(name string) {

}

func ToStateFunction(prod ebnf.Production) stateFn {

	// TODO Implement
	return func(lex *L, in LexemeGlob, ok chan<- match) stateFn {
		return stateFn(NewlineStateFn)
	}
}

/*
func NewGrammar(g ebnf.Grammar) *Grammar {
	grammar := &Grammar{}
	grammar.g = g
	lexemes := []string{}
	productions := []string{}

	for name := range g {
		first := []rune(name)[0]
		if unicode.IsLower(first) {
			lexemes = append(lexemes, name)
		} else {
			productions = append(productions, name)
		}
	}
	grammar.lexemes = lexemes
	grammar.productions = productions

	return grammar
}
*/
