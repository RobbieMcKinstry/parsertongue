package main

import (
	"unicode"
)

var (
	EOF = Token{
		typ:          eofType,
		val:          "EOF",
		FileLocation: FileLocation{-1, -1},
	}

	eofType = TokenType(0)

	prebuilt = map[string]stateFn{
		"newline":        runeMatcher(IsNewline),
		"unicode_char":   runeMatcher(IsUnicodeChar),
		"unicode_digit":  runeMatcher(IsUnicodeDigit),
		"unicode_letter": runeMatcher(unicode.IsLetter),
		"letter":         runeMatcher(IsLetter),
		"decimal_digit":  runeMatcher(IsDecimalDigit),
		"octal_digit":    runeMatcher(IsOctalDigit),
		"hex_digit":      runeMatcher(IsHexDigit),
	}
)
