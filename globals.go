package main

var (
	// TODO replace this with a genuine StartFn
	startFn = stateFn(NewlineStateFn)

	EOF = Token{
		typ:          eofType,
		val:          "EOF",
		FileLocation: FileLocation{-1, -1},
	}

	eofType = TokenType(0)
)
