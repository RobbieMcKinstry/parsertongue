package lexer

import "unicode"

type runeMatcher func(r rune) bool

// TODO do something about whitespace…
var prebuilt = map[string]runeMatcher{
	"newline":        runeMatcher(IsNewline),
	"unicode_char":   runeMatcher(IsUnicodeChar),
	"unicode_digit":  runeMatcher(IsUnicodeDigit),
	"unicode_letter": runeMatcher(unicode.IsLetter),
	"letter":         runeMatcher(IsLetter),
	"decimal_digit":  runeMatcher(IsDecimalDigit),
	"octal_digit":    runeMatcher(IsOctalDigit),
	"hex_digit":      runeMatcher(IsHexDigit),
	"whitespace":     runeMatcher(IsWhitespace),
}

// IsNewline checks if the rune is a newline char
func IsNewline(r rune) bool {
	return r == '\n'
}

// IsUnicodeChar checks if the rune is a unicode char
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

// IsWhitespace returns true if the rune is a
//Unicode whitespace character
func IsWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

// Whitespace is a stateFn which matches all
// whitespace in a row, returning the number of
// runes which are consecutive whitespace
func Whitespace(lex *L) (StateFn, int) {
	var count int
	for count = 0; IsWhitespace(lex.peek()); count++ {
		lex.next()
	}
	return nil, count
}
