package lexer

import (
	"testing"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
)

func TestPackageBlockOfficial(t *testing.T) {
	const root, path = "SourceFile", "../fixtures/package_block.ebnf"
	var sentence = []byte(`package main`)
	var gram = grammar.New(path, root)
	gram.Clean()
	var _, out = Lex(gram, sentence)
	var count = 0
	var expectedTokens = []string{"package", "main"}
	for token := range out {
		expected, observed := expectedTokens[count], token.Val
		if expected != observed {
			printTokenError(t, expected, observed)
		}
		count++
	}
	expected, observed := len(expectedTokens), count
	if expected != observed {
		t.Fatalf("Expected %v tokens but only found %v tokens", expected, observed)
	}
}

func TestImportBlock(t *testing.T) {
	t.Skip()
	const root, path = "S", "../fixtures/import_block.ebnf"
	var sentence = []byte(`
	package main
	
	import "hello"
	`)
	var gram = grammar.New(path, root)
	gram.Clean()
	var _, out = Lex(gram, sentence)
	var expectedTokens = []string{"package", "main", "import", "\"hello\""}
	var count = 0
	for token := range out {
		expected, observed := expectedTokens[count], token.Val
		if expected != observed {
			printTokenError(t, expected, observed)
		}
		count++
	}
	expected, observed := len(expectedTokens), count
	if expected != observed {
		t.Fatalf("Expected %v tokens but only found %v tokens", expected, observed)
	}
}
