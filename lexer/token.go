package lexer

import "golang.org/x/exp/ebnf"

// Token is a single lexeme
type Token struct {
	typ      *ebnf.Production
	Val      string
	IsLexeme bool
	FileLocation
}
