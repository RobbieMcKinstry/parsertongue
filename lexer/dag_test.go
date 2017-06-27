package lexer

import (
	"testing"

	"golang.org/x/exp/ebnf"
)

func TestTokenStateFnNoMatch(t *testing.T) {

	var (
		lex = buildLexer("hello world")
		tok = ebnf.Token{String: "golang"}
		fn  = lex.makeToken(&tok)
	)
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
		sentence = "golang"
		lex      = buildLexer(sentence)
		tok      = ebnf.Token{String: sentence}
		fn       = lex.makeToken(&tok)
	)
	nextFn, matchLen := fn(lex, 0)

	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := len(sentence), matchLen; expected != observed {
		t.Errorf("Expected no match: found %v", observed)
	}
}

func TestTokenStateFnPartialMatch(t *testing.T) {

	var (
		lex = buildLexer("go lang")
		tok = ebnf.Token{String: "golang"}
		fn  = lex.makeToken(&tok)
	)
	nextFn, matchLen := fn(lex, 0)

	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := 0, matchLen; expected != observed {
		t.Errorf("Expected no match: found %v", observed)
	}
}

func TestOptionalTrue(t *testing.T) {
	var lex = buildLexer("hello world")
	var tok = ebnf.Token{String: "hello"}
	var option = ebnf.Option{Body: &tok}
	var fn = lex.makeOption(&option)

	nextFn, matchLen := fn(lex, 0)
	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := len("hello"), matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}

func TestOptionalFalse(t *testing.T) {
	var lex = buildLexer("hell! world")
	var tok = ebnf.Token{String: "hello"}
	var option = ebnf.Option{Body: &tok}
	var fn = lex.makeOption(&option)

	nextFn, matchLen := fn(lex, 0)
	if observed := nextFn; observed != nil {
		t.Errorf("Expected %v, found %v", nil, observed)
	}

	if expected, observed := 0, matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}

func TestRepetitionTrue(t *testing.T) {
	var sentence = "go go go go "
	var lex = buildLexer(sentence)
	var tok = ebnf.Token{String: "go "}
	var rep = ebnf.Repetition{Body: &tok}
	var fn = lex.makeRepetition(&rep)

	matchLen := fn.Exhaust(lex, 0)

	if expected, observed := len(sentence), matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}
