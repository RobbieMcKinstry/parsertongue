package lexer

type stateFn func(lex *L, start int, ok chan<- match) stateFn
