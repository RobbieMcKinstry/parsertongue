package main

import (
	"unicode"
)

// ---------------------------------------------\\
// -------- PUT THIS IN ITS OWN FILE -----------\\
// ---------------------------------------------\\

// TODO implement
func NewlineStateFn(lex *L, start LexemeGlob, ok chan<- match) stateFn {
	return stateFn(NewlineStateFn)
}

// IsNewline checks if the rune is a newline char
func IsNewline(r rune) bool {
	return r == '\n'
}

// IsUnicodeChar checks if the rune is a unicode char
// TODO wtf is this doing??/
func IsUnicodeChar(r rune) bool {
	return !IsNewline(r)
}

// IsUnicodeDigit checks if the rune is a digit
func IsUnicodeDigit(r rune) bool {
	return unicode.IsDigit(r)
}

// IsLetter checks if the rune is a letter
func IsLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

// IsDecimalDigit checks if the rune is a decimal digit or not
func IsDecimalDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsOctalDigit checks if its an octal digit or not
func IsOctalDigit(r rune) bool {
	return '0' <= r && r <= '7'
}

// IsHexDigit checks if the rune is a hex digit or not.
func IsHexDigit(r rune) bool {
	return ('0' <= r && r <= '9') ||
		('A' <= r && r <= 'F') ||
		('a' <= r && r <= 'f')
}
