package grammar

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"unicode"

	"golang.org/x/exp/ebnf"
)

// G is a grammar
type G struct {
	// the underlying grammar
	gram ebnf.Grammar
	// lexemes is the list of Names of the strings
	// representing lexical productions
	lexemes []string

	// productions is the list of Names of the strings
	// representing syntactic productions
	prods []string

	// open is a mockable version of os.Open
	open func(string) (io.Reader, error)

	// SentenceRoot is the expected root production of this
	// grammar
	sentenceRoot string
}

// SentenceRoot will set the sentence root on this object
func (gram *G) SentenceRoot(root string) {
	gram.sentenceRoot = root
}

func wrapOpen() func(string) (io.Reader, error) {
	return func(path string) (io.Reader, error) {
		f, err := os.Open(path)
		return f, err
	}
}

// Grammar returns the grammar
func (gram *G) Grammar() ebnf.Grammar {
	return gram.gram
}

// LexemeNames returns the lex names
func (gram *G) LexemeNames() []string {
	return gram.lexemes
}

// ProdNames returns the prodnames
func (gram *G) ProdNames() []string {
	return gram.prods
}

// Prod returns the production with the
// given name
// TODO fix, this is like poisoning the well...
func (gram *G) Prod(name string) *ebnf.Production {
	prod, ok := gram.gram[name]
	if !ok {
		fmt.Println("Failed to receive the expected production")
		fmt.Printf("Prod is: %v\n", prod)
		fmt.Printf("Name is: %v\n", name)
		fmt.Printf("Gram is: %v\n", gram.gram)
		return nil
	}
	return prod
}

// IsLexeme returns true if the name
// starts with a lower case letter
func IsLexeme(name string) bool {
	if len(name) < 1 {
		return false
	}

	firstLetter := []rune(name)[0]
	return unicode.IsLower(firstLetter)
}

// Children returns the list of names reprenting the children
// of this tree node. It does not allow repeats (so it returns a set).
func Children(parent *ebnf.Production) []string {
	fmt.Println(Stringify(parent))
	if parent == nil {
		fmt.Println("Not supposed to receive a nil element...")
		return []string{}
	}
	altSet := parent.Expr
	return exprChildren(altSet)
}

// Stringify Turn a prod into a string
func Stringify(p *ebnf.Production) string {
	if p == nil {
		return "nil"
	}
	var rewriteRule = "nil"
	if p.Expr != nil {
		rewriteRule = ExprString(p.Expr)
	}

	return fmt.Sprintf("%v --> %v", p.Name.String, rewriteRule)
}

// ExprString is a toString method for an expression type
func ExprString(exp ebnf.Expression) string {
	var res string
	switch v := exp.(type) {
	case *ebnf.Option:
		subExpr := ExprString(v.Body)
		res = fmt.Sprintf("[ %v ]", subExpr)
	case *ebnf.Group:
		subExpr := ExprString(v.Body)
		res = fmt.Sprintf("( %v )", subExpr)
	case *ebnf.Repetition:
		subExpr := ExprString(v.Body)
		res = fmt.Sprintf("{ %v }", subExpr)
	case ebnf.Sequence:
		var subExpr string
		for i, alt := range v {
			if i > 0 {
				subExpr += " | "
			}
			fmt.Println("prod", alt)
			subExpr += ExprString(alt)
		}
		res = subExpr

	case ebnf.Alternative:
		var subExpr string
		for i, alt := range v {
			if i > 0 {
				subExpr += " | "
			}
			subExpr += ExprString(alt)
		}
		res = subExpr

	case *ebnf.Name:
		res = v.String
	case *ebnf.Token:
		res = fmt.Sprintf("\"%v\"", v.String)
	default:
		res = "error"
		fmt.Println(reflect.TypeOf(v))
	}
	return res
}

// setAppend will append elem to the slice
// iff elem is not already in the slice. Otherwise, it
// functions just like the built-in append.
func setAppend(in []string, elems ...string) (out []string) {
	out = in
	for _, elem := range elems {
		if !strContains(out, elem) {
			out = append(out, elem)
		}
	}
	return out
}

// strContains returns true if the string elem is contained
// in the slice
func strContains(slice []string, elem string) bool {
	var res = false
	for _, contained := range slice {
		if elem == contained {
			res = true
			break
		}
	}
	return res
}

func exprChildren(parent ebnf.Expression) []string {

	var children = []string{}
	switch v := parent.(type) {
	case ebnf.Alternative:
		seq := []ebnf.Expression(v)
		for _, expr := range seq {
			child := exprChildren(expr)
			children = setAppend(children, child...)
		}
		return children
	case ebnf.Sequence:
		seq := []ebnf.Expression(v)
		for _, expr := range seq {
			child := exprChildren(expr)
			children = setAppend(children, child...)
		}
		return children
	case *ebnf.Repetition:
		return exprChildren(v.Body)
	case *ebnf.Option:
		return exprChildren(v.Body)
	case *ebnf.Group:
		return exprChildren(v.Body)
	case *ebnf.Name:
		return []string{v.String}

	// Ranges are all literals, and thus should not
	// be counted as children, since they are ternimal
	// symbols
	case *ebnf.Range:
		return []string{}
	}
	return nil
}

func newG() *G {
	return &G{
		open: wrapOpen(),
	}
}

// New makes a Grammar
func New(path, root string) *G {
	g := newG()
	g.SentenceRoot(root)
	g.init(path)
	return g
}

func (gram *G) init(path string) {

	file, err := gram.open(path)
	if err != nil {
		panic(err)
	}

	grammar, err := ebnf.Parse(path, file)
	if err != nil {
		panic(err)
	}
	verifyGrammar(grammar, gram.sentenceRoot)
	gram.gram = grammar
	gram.lexemes = lowercaseProds(grammar)
	gram.prods = uppercaseProds(grammar)
}

func lowercaseProds(gram ebnf.Grammar) []string {
	if gram == nil {
		return []string{}
	}

	var res []string
	for name := range gram {
		if IsLexeme(name) {
			res = append(res, name)
		}
	}
	return res
}

func uppercaseProds(gram ebnf.Grammar) []string {
	if gram == nil {
		return []string{}
	}

	var res []string
	for name := range gram {
		if !IsLexeme(name) {
			res = append(res, name)
		}
	}
	return res
}

// VerifyGrammar checks to make sure that the
// grammar is ebnf.Verified
func verifyGrammar(g ebnf.Grammar, root string) ebnf.Grammar {
	if err := ebnf.Verify(g, root); err != nil {
		panic(err)
	}
	return g
}

// FindEntrantProds returns the lexical productions which are entered
// by non-lexical productions
func FindEntrantProds(gram *G) []string {
	var (
		entrantSet = map[string]bool{}
		names      = gram.ProdNames()
	)

	for _, name := range names {
		// fetch the production out of the grammar
		var prod *ebnf.Production = gram.Prod(name)
		// Now, walk the production and get it's children...
		children := Children(prod)
		for _, child := range children {
			if IsLexeme(child) {
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
