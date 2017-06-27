package lexer

func buildLexer(sentence string) *L {
	lex := new(L)
	data := []byte(sentence)
	lex.reader = NewBufferScanner(data)
	return lex
}
