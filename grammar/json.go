package grammar

import (
	"golang.org/x/exp/ebnf"
)

// JSON is the struct which returns JSON
// representation of the *grammar.G
type JSON struct {
	Root    string        `json:"root"`
	Grammar *ebnf.Grammar `json:"grammar"`
}

// NewJSON returns a new JSON Obj
func NewJSON(gram *G) *JSON {
	return &JSON{
		Root:    gram.sentenceRoot,
		Grammar: &gram.gram,
	}
}

// NewJSONFile builds the file from the JSON
func NewJSONFile(filename, root string) *JSON {
	g := New(filename, root)
	return &JSON{
		Root:    root,
		Grammar: &g.gram,
	}
}
