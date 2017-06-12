package lexer

import "fmt"

type FileLocation struct {
	Start, End int
}

func (fl FileLocation) String() string {
	return fmt.Sprintf("(%i, %i)")
}
