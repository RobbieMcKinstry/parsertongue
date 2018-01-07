package lexer

import (
	"fmt"
	"testing"

	"github.com/RobbieMcKinstry/parsertongue/grammar"
	"golang.org/x/exp/ebnf"
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

func TestLiteralExample1(t *testing.T) {
	const root, path = "S", "../fixtures/string_literal.ebnf"
	var stringLit = `"this is not hello world"`
	var sentence = []byte(stringLit)
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	fmt.Println("Testing literal example.")
	for val := range out {
		if expected, observed := val.Val, stringLit; expected != observed {
			t.Fatalf("Expected %v but found %v", expected, observed)
		}
	}
}

func TestLiteral2Example1(t *testing.T) {
	const root, path = "S", "../fixtures/string_literal2.ebnf"
	var stringLit = `"this is not hello world"`
	var sentence = []byte(stringLit)
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	for val := range out {
		if expected, observed := val.Val, stringLit; expected != observed {
			t.Fatalf("Expected %v but found %v", expected, observed)
		}
	}
}

func TestLiteral2Example2(t *testing.T) {
	const root, path = "S", "../fixtures/string_literal2.ebnf"
	var stringLit = `"\U200B this is not \"hello world\""`
	var sentence = []byte(stringLit)
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	for val := range out {
		if expected, observed := val.Val, stringLit; expected != observed {
			t.Fatalf("Expected %v but found %v", expected, observed)
		}
	}
}

func TestRun1Example1(t *testing.T) {
	const root, path = "S", "../fixtures/01.ebnf"
	var sentence = []byte("foo")
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	var count = 0
	for token := range out {
		if count++; count >= 2 {
			t.Error("Too many tokens received!")
		}
		if token.Val != "foo" {
			t.Errorf("Expected to receive the string 'foo' but found %v", token.Val)
		}
		if token.typ != nil {
			t.Errorf("Expected a nil type but found a value")
		}
		if !token.IsLexemeLiteral {
			t.Errorf("Expected IsLexemeLiteral to be true.")
		}
	}
}

func TestRun1Example2(t *testing.T) {
	const root, path = "S", "../fixtures/01.ebnf"
	var sentence = []byte("hello")
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	var count = 0
	for token := range out {
		if count++; count >= 2 {
			t.Error("Too many tokens received!")
		}
		if token.Val != "hello" {
			t.Errorf("Expected to receive the string 'hello' but found %v", token.Val)
		}
		if token.typ != gram.Prod("hello") {
			t.Errorf(
				"Expected to receive the production 'S', found '%v'",
				token.typ.Name.String,
			)
		}
	}
}

func TestRun1Example3(t *testing.T) {
	const root, path = "S", "../fixtures/01.ebnf"
	var sentence = []byte("failure")
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	for token := range out {
		t.Errorf("Expected no tokens, found %v", token)
	}
}

func TestRun2Example1(t *testing.T) {
	const root, path = "S", "../fixtures/02.ebnf"
	var sentence = []byte("goodbye my love!")
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	var expectedTokens = []string{"goodbye", "my love", "!"}
	var expectedProductions = []*ebnf.Production{
		gram.Prod("foo"), gram.Prod("bar"), gram.Prod("baz"),
	}
	var count = 0
	for token := range out {
		if count >= 4 {
			t.Error("Too many tokens received!")
		}
		if expected, observed := expectedTokens[count], token.Val; expected != observed {
			printTokenError(t, expected, observed)
		}
		if expected, observed := expectedProductions[count], token.typ; expected != observed {
			printProductionError(t, expected, observed)
		}

		count++
	}
}

func TestSimpleGolang(t *testing.T) {
	const root, path = "SourceFile", "../fixtures/golang_augmented.ebnf"
	var sentence = []byte(`package main

	var x int`)

	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	var count = 0
	var expectedTokens = []string{"package", "main", "var", "x", "int"}
	for token := range out {
		fmt.Println(token.Val)
		if expected, observed := expectedTokens[count], token.Val; expected != observed {
			printTokenError(t, expected, observed)
		}
		count++
	}
	if expected, observed := len(expectedTokens), count; expected != observed {
		t.Fatalf("Expected %v tokens but only found %v tokens", expected, observed)
	}
}

func TestSimpleGolang2(t *testing.T) {
	t.Skip()
	const root, path = "SourceFile", "../fixtures/golang_augmented.ebnf"
	var sentence = []byte(`package main
	import "fmt"

	func main() {
		fmt.Println("hello world")
	}`)

	var gram = grammar.New(path, root)
	// gram.Clean()
	var _, out = Lex(gram, sentence)
	var count = 0
	var expectedTokens = []string{"package", "main", "import",
		"fmt", "func", "main", "(", ")", "{", "fmt", ".", "Println", "(", "hello world", ")", "}"}
	for token := range out {
		fmt.Println(token.Val)
		if expected, observed := expectedTokens[count], token.Val; expected != observed {
			printTokenError(t, expected, observed)
		}
		count++
	}
	if expected, observed := len(expectedTokens), count; expected != observed {

		// fmt.Printf("Here's the grammar:\n%v", gram)
		t.Fatalf("Expected %v tokens but only found %v tokens", expected, observed)
	}
}

func printTokenError(t *testing.T, expected, observed string) {
	t.Fatalf("Expected token \"%v\" but found \"%v\"", expected, observed)
}

func printProductionError(t *testing.T, expected, observed *ebnf.Production) {
	t.Fatalf(
		"Expected to receive the production '%v', found '%v'",
		expected.Name, observed.Name,
	)
}
