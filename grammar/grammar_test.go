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

func TestEntrantProds01(t *testing.T) {
	const root, path = "S", "../fixtures/01.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 1, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}

	if ok := containsString(entrants, "hello"); !ok {
		t.Errorf("Failed to find entrant production \"%v\"", "hello")
	}
}

func TestEntrantProds02(t *testing.T) {
	const root, path = "S", "../fixtures/02.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 3, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}

	entrantsExpected := []string{"foo", "bar", "baz"}
	for _, expected := range entrantsExpected {
		if ok := containsString(entrants, expected); !ok {
			t.Errorf("Failed to find entrant production \"%v\"",
				expected,
			)
		}
	}
}

func TestEntrantProds03(t *testing.T) {
	const root, path = "N", "../fixtures/03.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 1, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}

	entrant := "empty"
	if ok := containsString(entrants, entrant); !ok {
		t.Errorf("Failed to find entrant production \"%v\"",
			entrant,
		)
	}
}

func TestEntrantProds04(t *testing.T) {
	const root, path = "Integer", "../fixtures/04.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 1, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}

	entrant := "empty"
	if ok := containsString(entrants, entrant); !ok {
		t.Errorf("Failed to find entrant production \"%v\"",
			entrant,
		)
	}
}

func TestEntrantProds05(t *testing.T) {
	const root, path = "Party", "../fixtures/05.ebnf"
	var gram = New(path, root)
	entrants := FindEntrantProds(gram)

	if expected, observed := 1, len(entrants); expected != observed {
		t.Fatalf("Incorrect number of entrant prods: expected %v, found %v",
			expected,
			observed,
		)
	}
	entrant := "bard"
	if ok := containsString(entrants, entrant); !ok {
		t.Errorf("Failed to find entrant production \"%v\"",
			entrant,
		)
	}

}

func TestChildren01(t *testing.T) {
	const root, path = "S", "../fixtures/01.ebnf"

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

func TestChildren02(t *testing.T) {
	const root, path = "S", "../fixtures/02.ebnf"

	var (
		gram     = New(path, root)
		rootRule = gram.Prod(root)
		children = Children(rootRule)
	)

	if expected, observed := 3, len(children); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "foo", children[0]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "bar", children[1]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "baz", children[2]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}
}

func TestChildren03(t *testing.T) {
	const root, path = "N", "../fixtures/03.ebnf"

	var (
		gram     = New(path, root)
		rootRule = gram.Prod(root)
		children = Children(rootRule)
	)

	if expected, observed := 2, len(children); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Number", children[0]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "empty", children[1]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}
}

func TestChildren04(t *testing.T) {
	const root, path = "Integer", "../fixtures/04.ebnf"

	var (
		gram          = New(path, root)
		rootRule      = gram.Prod(root)
		digitRule     = gram.Prod("Digit")
		children      = Children(rootRule)
		rangeChildren = Children(digitRule)
	)

	if expected, observed := 3, len(children); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Digit", children[0]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "empty", children[1]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Integer", children[2]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := 0, len(rangeChildren); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}
}

func TestChildren05(t *testing.T) {
	const root, path = "Party", "../fixtures/05.ebnf"

	var (
		gram     = New(path, root)
		rootRule = gram.Prod(root)
		children = Children(rootRule)
	)

	if expected, observed := 4, len(children); expected != observed {
		t.Errorf("Incorrect length: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Fighter", children[0]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Mage", children[1]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "Ranger", children[2]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}

	if expected, observed := "bard", children[3]; expected != observed {
		t.Errorf("Incorrect value: expected %v, found %v",
			expected,
			observed,
		)
	}
}

func TestSetAppend(t *testing.T) {
	var table = []struct {
		in       []string
		arg      string
		expected []string
	}{
		{
			in:       []string{"foo", "bar"},
			arg:      "baz",
			expected: []string{"foo", "bar", "baz"},
		}, {
			in:       []string{},
			arg:      "single",
			expected: []string{"single"},
		}, {
			in:       []string{"abcd"},
			arg:      "abcd",
			expected: []string{"abcd"},
		}, {
			in:       []string{"a", "b", "c", "d"},
			arg:      "c",
			expected: []string{"a", "b", "c", "d"},
		},
	}

	for _, row := range table {
		observed := setAppend(row.in, row.arg)
		matchStringSets(t, observed, row.expected)
	}
}

func matchStringSets(t *testing.T, x, y []string) {
	if x == nil || y == nil {
		t.Fatal("Nil slice")
	}

	if len(x) != len(y) {
		t.Errorf("Mismatched lengths: %v != %v", len(x), len(y))
	}
}
