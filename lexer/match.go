package lexer

import "fmt"

type match struct {
	start, width int
}

// String converts the match into a string
func (m match) String() string {
	return fmt.Sprintf("[%d:%d]", m.start, m.start+m.width)
}

func (m match) Start() int {
	return m.start
}

// Returns the exclusive end of the match.
func (m match) End() int {
	return m.start + m.width - 1
}

func (m match) Width() int {
	return m.width
}
