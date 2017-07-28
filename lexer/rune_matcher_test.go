package lexer

import "testing"

func TestDigitMatcher(t *testing.T) {
	lex := new(L)
	r := prebuilt["unicode_digit"]
	fn := lex.makeRuneMatcher(r)
	sentence := []byte("0123456789")
	lex.reader = NewBufferScanner(sentence)

	for range sentence {
		nextFn, matchLen := fn(lex.Clone())
		if observed := nextFn; observed != nil {
			t.Errorf("Expected %v, found %v", nil, observed)
		}

		if expected, observed := 1, matchLen; expected != observed {
			t.Errorf("Expected match length %v, found %v", expected, observed)
		}

		lex.next()
	}

	finalRune := lex.next()
	if expected, observed := eof, finalRune; expected != observed {
		t.Errorf("Mismatched final runes: %c != %c", expected, observed)
	}
}

func TestDigitMatcherLetters(t *testing.T) {
	lex := new(L)
	r := prebuilt["unicode_digit"]
	fn := lex.makeRuneMatcher(r)
	sentence := []byte("abcdefghijk")
	lex.reader = NewBufferScanner(sentence)

	for range sentence {
		nextFn, matchLen := fn(lex.Clone())
		if observed := nextFn; observed != nil {
			t.Errorf("Expected %v, found %v", nil, observed)
		}

		if expected, observed := -1, matchLen; expected != observed {
			t.Errorf("Expected match length %v, found %v", expected, observed)
		}

		lex.next()
	}

	finalRune := lex.next()
	if expected, observed := eof, finalRune; expected != observed {
		t.Errorf("Mismatched final runes: %c != %c", expected, observed)
	}
}
