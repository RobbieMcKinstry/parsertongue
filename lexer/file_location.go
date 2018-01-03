package lexer

import "fmt"

// FileLocation stores the location in the file, indexed in by number of bytes
type FileLocation struct {
	Start, End int
}

// String prints the string representation for the user.
func (fl FileLocation) String() string {
	return fmt.Sprintf("(%d, %d)", fl.Start, fl.End)
}
