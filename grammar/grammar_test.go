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
		t.Error("Expected %v uppercase prod but found %v", expectedLen, length)
	}

	foundProd := uppers[0]
	if foundProd != "Uppercase" {
		t.Error("Expected %v to be %v", foundProd, prod2)
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
		t.Error("Expected %v uppercase prod but found %v", expectedLen, length)
	}

	foundProd := lowers[0]
	if foundProd != "lowercase" {
		t.Error("Expected %v to be %v", foundProd, prod2)
	}
}
