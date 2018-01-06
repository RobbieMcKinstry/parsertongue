package grammar

import (
	"testing"

	"golang.org/x/exp/ebnf"
)

func TestCleanExample1(t *testing.T) {
	// Read in a sample grammar.
	const root, path = "S", "../fixtures/01.ebnf"
	var gram = New(path, root)
	gram.clean()

	// now I should expect to have a production
	// named "\x200B_lexeme_foo".
	name := "\u200B_lexeme_foo"

	found := gram.Prod(name)
	if found == nil {
		t.Fatalf("Expected non-nil production, found nil")
	}
	if expected, observed := name, found.Name.String; expected != observed {
		t.Fatalf("Expected name \"%v\", found \"%v\"", expected, observed)
	}
}

func TestClean2Example1(t *testing.T) {
	// Read in a sample grammar.
	const root, path = "S", "../fixtures/01.ebnf"
	var gram = New(path, root)
	gram.clean()

	// now I should NOT expect to have a production
	// named "\x200B_lexeme_hello".
	name := "\u200B_lexeme_hello"

	found := gram.Prod(name)
	if found != nil {
		t.Fatalf("Expected nil production, found %v\nGrammar is:\n%v", Stringify(found), gram)
	}
}

func TestLiteralNameFromLexeme(t *testing.T) {
	token := &ebnf.Token{
		String: "=",
	}
	expected, observed := "\u200B_lexeme_=", literalNameFromLexeme(token)
	if expected != observed {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, observed)
	}
}

func TestLiteralNameFromLexeme2(t *testing.T) {
	token := &ebnf.Token{
		String: "foo",
	}
	expected, observed := "\u200B_lexeme_foo", literalNameFromLexeme(token)
	if expected != observed {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, observed)
	}
}

func TestLiteralNameFromLexeme3(t *testing.T) {
	token := &ebnf.Token{
		String: "barrrrrr",
	}
	expected, observed := "\u200B_lexeme_wrong_value", literalNameFromLexeme(token)
	if expected == observed {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, observed)
	}
}

func TestLexemeNameFromLiteral(t *testing.T) {
	origName := "foo"
	token := &ebnf.Token{
		String: origName,
	}
	newName := literalNameFromLexeme(token)
	token.String = newName
	foundName := lexemeNameFromLiteral(token)

	expected, observed := origName, foundName
	if expected != observed {
		t.Fatalf("Expected \"%v\" but found \"%v\"", expected, observed)
	}
}
