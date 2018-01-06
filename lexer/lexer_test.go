package lexer

import (
	"fmt"
	"io/ioutil"
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
	var count = 0
	for token := range out {
		if count++; count >= 4 {
			t.Error("Too many tokens received!")
		}
		switch count {
		case 1:
			if token.Val != "goodbye" {
				t.Errorf(
					"Expected \"goodbye\", found \"%v\"",
					token.Val,
				)
			}
			if token.typ != gram.Prod("foo") {
				t.Errorf(
					"Expected to receive the production 'foo', found '%v'",
					token.typ.Name.String,
				)
			}

		case 2:
			if token.Val != "my love" {
				t.Errorf(
					"Expected \"my love\", found \"%v\"",
					token.Val,
				)
			}
			if token.typ != gram.Prod("bar") {
				t.Errorf(
					"Expected to receive the production 'bar', found '%v'",
					token.typ.Name.String,
				)
			}

		case 3:
			if token.Val != "!" {
				t.Errorf(
					"Expected \"!\", found \"%v\"",
					token.Val,
				)
			}
			if token.typ != gram.Prod("baz") {
				t.Errorf(
					"Expected to receive the production 'baz', found '%v'",
					token.typ.Name.String,
				)
			}
		}
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

func TestRunGolang(t *testing.T) {
	t.Skip()
	const root, path = "SourceFile", "../fixtures/golang_semi.ebnf"
	var sentence, err = ioutil.ReadFile("lexer_test.go")
	if err != nil {
		t.Fatal(err)
	}
	var gram = grammar.New(path, root)
	var _, out = Lex(gram, sentence)
	var count = 0
	for token := range out {
		count++
		fmt.Println(token.Val)
	}
	if count < 200 {
		t.Fatal("There are definitely more than 200 tokens in this file.")
	}
}
