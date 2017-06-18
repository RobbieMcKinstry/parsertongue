package ll1

import (
	"github.com/RobbieMcKinstry/parsertongue/grammar"
	"golang.org/x/exp/ebnf"
)

func findEntrantProds(gram *grammar.G) {
	var (
		entrantSet = map[string]bool{}
		names      = gram.ProdNames()
	)

	for _, name := range names {
		// fetch the production out of the grammar
		var prod *ebnf.Production = gram.Prod(name)
		// Now, walk the production and get it's children...
		children := grammar.Children(prod)
		for _, child := range children {
			if grammar.IsLexeme(child) {
				// add it to the entrantSet
				entrantSet[child] = true
			}
		}
	}

	// convert the map into a list
	var entrantProds = make([]string, 0, len(entrantSet))
	for name := range entrantSet {
		entrantProds = append(entrantProds, name)
	}

	return entrantProds
}
