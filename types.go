package main

import (
	"golang.org/x/exp/ebnf"
)

type (
	L struct {
		g      ebnf.Grammar
		input  string
		tokens chan<- Token // channel of scanned tokens
	}

	Token struct {
		typ TokenType
		val string
		FileLocation
	}

	LexemeGlob struct {
		start, end, width int
		file              *string
	}

	FileLocation struct {
		start, end int
	}

	TokenType int

	Grammar struct {
		// store the EBNF grammar
		g ebnf.Grammar

		lexemes     []string
		productions []string
	}

	stateFn func(lex *L, in LexemeGlob, ok chan<- match) stateFn

	match struct {
		start, end LexemeGlob
	}
)
