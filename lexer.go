package main

import (
	"io"
	"io/ioutil"
	"log"

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

// TODO move this into a global file
var startFn = stateFn(NewlineStateFn)

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
