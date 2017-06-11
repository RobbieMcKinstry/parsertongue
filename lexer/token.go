package lexer

type Token struct {
	typ TokenType
	val string
	FileLocation
}
