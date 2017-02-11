package main

import (
	"fmt"
)

func (t Token) String() string {
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

var EOF = Token{
	typ:          eofType,
	val:          "EOF",
	FileLocation: FileLocation{-1, -1},
}

var eofType = TokenType(0)
