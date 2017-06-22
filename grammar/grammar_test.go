package grammar

import (
	"testing"

	"golang.org/x/exp/ebnf"
)

func TestUppercaseProds(t *testing.T) {
	// make a new grammar...
	var (
		gram                  = make(ebnf.Grammar)
		alt1 ebnf.Alternative = []ebnf.Expression{}
		alt2 ebnf.Alternative = []ebnf.Expression{}

		prod1 = ebnf.Production{Expr: alt1}
		prod2 = ebnf.Production{Expr: alt2}
	)
	gram["lowercase"] = &prod1
	gram["Uppercase"] = &prod2

	uppers := uppercaseProds(gram)
	length := len(uppers)

	const expectedLen = 1

	if length != expectedLen {
		t.Errorf("Expected %v uppercase prod but found %v", expectedLen, length)
	}

	foundProd := uppers[0]
	if foundProd != "Uppercase" {
		t.Errorf("Expected %v to be %v", foundProd, prod2)
	}
}

func TestLowercaseProds(t *testing.T) {
	// make a new grammar...
	var (
		gram                  = make(ebnf.Grammar)
		alt1 ebnf.Alternative = []ebnf.Expression{}
		alt2 ebnf.Alternative = []ebnf.Expression{}

		prod1 = ebnf.Production{Expr: alt1}
		prod2 = ebnf.Production{Expr: alt2}
	)
	gram["lowercase"] = &prod1
	gram["Uppercase"] = &prod2

	lowers := lowercaseProds(gram)
	length := len(lowers)

	const expectedLen = 1

	if length != expectedLen {
		t.Errorf("Expected %v uppercase prod but found %v", expectedLen, length)
	}

	foundProd := lowers[0]
	if foundProd != "lowercase" {
		t.Errorf("Expected %v to be %v", foundProd, prod2)
	}
}

func TestEntrantProds(t *testing.T) {
	// TODO Fix
	// test the "FindEntrantProds" function
	// t.Skip()
	const root, path = "S", "test_fixtures/01.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 1, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "hello", entrants[0]; expected != observed {
		t.Errorf("Incorrect entrant production: expected %v, found %v",
			expected,
			observed,
		)
	}
}

func TestChildren(t *testing.T) {
	const root, path = "S", "test_fixtures/01.ebnf"

	var (
		gram     = New(path, root)
		rootRule = gram.Prod(root)
		children = Children(rootRule)
	)

	if expected, observed := 1, len(children); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "hello", children[0]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)

	}
}
