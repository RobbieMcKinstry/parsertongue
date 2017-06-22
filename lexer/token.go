package lexer

type Token struct {
	typ TokenType
	val string
	FileLocation
}

type TokenType int

const (
	Keyword TokenType = iota
	ID
	StringLit
	NumericLit
	Operator
)
