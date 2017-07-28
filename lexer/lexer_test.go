package lexer

import (
	"testing"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
)

func TestCollectProds(t *testing.T) {
	const root, path = "Integer", "../fixtures/04.ebnf"
	const sentence = ""
	var gram = grammar.New(path, root)
	var lex = buildLexer(sentence)
	var names = []string{"Digit", "Integer", "empty"}
	lex.gram = gram
	var prods = lex.collectProdsByName(names)

	for i, prod := range prods {
		name := prod.Name.String
		if name != names[i] {
			t.Errorf("Name mismatch: %v != %v", name, names[i])
		}
	}
}

func TestRun1(t *testing.T) {
	t.Skip()
	const root, path = "S", "../fixtures/01.ebnf"
	var sentence = []byte("foo")
	var gram = grammar.New(path, root)
	var _, out = lex(gram, sentence)
	var count = 0
	for token := range out {
		if count++; count >= 2 {
			t.Error("Too many tokens received!")
		}
		if token.val != "foo" {
			t.Errorf("Expected to receive the string 'foo' but found %v", token.val)
		}
		if token.typ != gram.Prod("S") {
			t.Error("Expected to receive the production 'S'")
		}
	}
}
