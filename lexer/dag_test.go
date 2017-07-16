package lexer

import (
	"testing"

	"github.com/RobbieMcKinstry/parsertongue/grammar"

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

func TestPrebuild1(t *testing.T) {
	var sentence = "\n"
	var lex = buildLexer(sentence)
	var newline = ebnf.Name{String: "newline"}
	var fn = lex.makeName(&newline)

	matchLen := fn.Exhaust(lex, 0)
	if expected, observed := len(sentence), matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}

func TestPrebuild2(t *testing.T) {
	var sentence = "foo"
	var lex = buildLexer(sentence)
	var newline = ebnf.Name{String: "newline"}
	var fn = lex.makeName(&newline)

	matchLen := fn.Exhaust(lex, 0)
	if expected, observed := 0, matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}

// Smash my face against the keyboard to generate a random
// Unicode string.
func TestPrebuild3(t *testing.T) {
	var sentence = "foo123“…æå∂ƒ©ƒ®´†´ß√˜∫†˙µ…¬†"
	var lex = buildLexer(sentence)
	var unicodeChar = ebnf.Name{String: "unicode_char"}
	var fn = lex.makeName(&unicodeChar)

	for _, character := range sentence {
		matchLen := fn.Exhaust(lex.Clone(), 0)
		if expected, observed := 1, matchLen; expected != observed {
			t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
		}
		r := lex.next()
		if r != character {
			t.Errorf("Mismatched characters: %v != %v", r, character)
		}
	}
}

func TestMakeName(t *testing.T) {
	// Grab an existing grammar
	const (
		root, path = "Party", "../fixtures/05.ebnf"
		fighter    = "fighter"
		mage       = "mage"
		sentence   = "fighter mage"
	)

	var (
		gram        = grammar.New(path, root)
		lex         = buildLexer(sentence)
		fighterExpr = ebnf.Name{String: "Fighter"}
		mageExpr    = ebnf.Name{String: "Mage"}
	)
	lex.gram = gram
	var (
		fnFighter = lex.makeName(&fighterExpr)
		fnMage    = lex.makeName(&mageExpr)
		matchLen  = fnFighter.Exhaust(lex.Clone(), 0)
	)

	if expected, observed := len(fighter), matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}

	// advance the lexer to the next production
	for range fighter {
		lex.next()
	}
	// consume the whitespace
	lex.next()

	// now, match the next word…
	matchLen = fnMage.Exhaust(lex.Clone(), 0)
	if expected, observed := len(mage), matchLen; expected != observed {
		t.Errorf("Unexpected Match length: expected %v but found %v", expected, observed)
	}
}
