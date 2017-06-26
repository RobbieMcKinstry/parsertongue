package lexer

import (
	"testing"

	"golang.org/x/exp/ebnf"
)

func TestTokenStateFnNoMatch(t *testing.T) {

	var (
		lex      = new(L)
		sentence = []byte("hello world")
		tok      = ebnf.Token{String: "golang"}
		fn       = lex.makeToken(&tok)
	)
	lex.reader = NewBufferScanner(sentence)
	nextFn, matchLen := fn(lex, 0)

	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := 0, matchLen; expected != observed {
		t.Errorf("Expected no match: found %v", observed)
	}
}

func TestTokenStateFnYesMatch(t *testing.T) {

	var (
		lex      = new(L)
		sentence = []byte("golang")
		tok      = ebnf.Token{String: "golang"}
		fn       = lex.makeToken(&tok)
	)
	lex.reader = NewBufferScanner(sentence)
	nextFn, matchLen := fn(lex, 0)

	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := len(sentence), matchLen; expected != observed {
		t.Errorf("Expected no match: found %v", observed)
	}
}
