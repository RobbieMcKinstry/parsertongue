package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/exp/ebnf"
)

func NewLexer(g ebnf.Grammar, r io.Reader, tokenStream chan<- Token) *L {
	payload, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	return &L{
		g:      g,
		input:  string(payload),
		tokens: tokenStream,
	}
}

func (lex *L) emit(t Token) {
	lex.tokens <- t
}

func (lex *L) run() {
	// Start at the start state -> this should be at the very top of the tree.

	lexGlob := LexemeGlob{0, 0, 0, &lex.input}

	// this is NOT right.
	// We shouldn't be spinning endlessly here, I don't think.
	// We should definitely be more careful about how this works and what it's returning.
	for state := startFn; state != nil; state = state(lex, lexGlob, make(chan match)) {
	}
	close(lex.tokens)
}

func lex(g ebnf.Grammar, f *os.File) (*L, chan<- Token) {

	channel := make(chan Token)
	lexer := NewLexer(g, f, channel)
	go lexer.run()
	return lexer, channel
}
