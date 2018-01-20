package grammar

// JSON is the struct which returns JSON
// representation of the *grammar.G
type JSON struct {
	Root        string              `json:"root"`
	Productions map[string][]string `json:"grammar"`
}

// NewJSON returns a new JSON Obj
func NewJSON(gram *G) *JSON {
	jsonObj := &JSON{
		Root:        gram.sentenceRoot,
		Productions: map[string][]string{},
	}
	for name, prod := range gram.gram {
		jsonObj.Productions[name] = Children(prod)
	}
	return jsonObj
}

// NewJSONFile builds the file from the JSON
func NewJSONFile(filename, root string) *JSON {
	g := New(filename, root)
	return NewJSON(g)
}
