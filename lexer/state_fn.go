package lexer

// StateFn represents a node in the the NFA.
type StateFn func(lex *L, start int) (StateFn, int)
