package grammar

import (
	"fmt"
	"strings"

	"golang.org/x/exp/ebnf"
)

// Clean iterates through the grammar, looking for all
// token literals. It replaces all token literals with
// productions which point to a single token. It adds
// those productions to the grammar.
// The produced productions are all lexical, since all
// token literals are entrant by definition.
func (gram *G) Clean() {
	// Collect the tokens
	tokenMap := gram.FindTokenLiterals()

	// Now, for each of those tokens,
	// create a new production that contains only that token literal
	var productions = lexemeProds(tokenMap)

	// Replace the tokens in the non-lexical productions
	// with a reference to the newly created lexical production
	// This step must be performed before the next otherwise
	// this will cause each of the newly added productions to
	// be added twice.
	gram.replaceTokenLiterals(tokenMap)

	// finally, add each of the lexical productions to the grammar:
	gram.addProds(productions)
}

// addProds takes a mapping from the name of the production to the
// productio struct and adds them to the grammar.
func (gram *G) addProds(prods map[string]*ebnf.Production) {
	for name, prod := range prods {
		// check to see of the existing production is already
		// at the top level.
		gram.gram[name] = prod
	}
}

// This function builds the lexical productions using the entrant token
// literals found in the non-lexical productions.
func lexemeProds(toks map[string]*ebnf.Token) map[string]*ebnf.Production {
	var productions = map[string]*ebnf.Production{}
	for _, tok := range toks {
		prod := marshalIntoProduction(tok)
		productions[prod.Name.String] = prod
	}
	return productions
}

// marshalIntoProduction builds the lexical production
// from the token.
func marshalIntoProduction(tok *ebnf.Token) *ebnf.Production {
	// now, prepend a zero-width whitespace character.
	edited := literalNameFromLexeme(tok)

	// ok, now slam this sucker into a production:
	prod := &ebnf.Production{
		Name: &ebnf.Name{
			String: edited,
		},
		Expr: tok,
	}
	return prod
}

// LiteralNameFromLexeme takes a lexeme and prepends the
// magic prefix.
func literalNameFromLexeme(tok *ebnf.Token) string {
	body := tok.String
	edited := fmt.Sprintf("\u200B_lexeme_%v", body)
	return edited
}

// lexemeNameFromLiteral takes a token with a magic
// prefix and strips that prefix.
func lexemeNameFromLiteral(tok *ebnf.Token) string {
	body := tok.String
	// Chop off the first couple of characters.
	edited := strings.Replace(body, "\u200B_lexeme_", "", 1)
	return edited
}

// replaceTokenLiterals calls down to a bunch of functions…
// it does the act of replacing each of the token in
// non-lexical productions
func (gram *G) replaceTokenLiterals(oldToks map[string]*ebnf.Token) {
	for name, production := range gram.gram {
		// dont replace tokens in lexemes
		if IsLexeme(name) {
			continue
		}
		// check special cases:
		// 1. expr == nil
		// 2. typeof(expr) == token
		if productionIsTokenSingleton(production) {
			continue
		}
		gram.replaceExpressionIfLiteral(production.Expr)
	}
}

// returns true if the expression is nil or if
// the expression is a token only.
func productionIsTokenSingleton(prod *ebnf.Production) bool {
	expr := prod.Expr
	if expr == nil {
		return true
	}
	if _, ok := expr.(*ebnf.Token); ok {
		return true
	}
	return false
}

// nameReferenceForToken converts the given token
// into an *ebnf.Name which would reference it, should it be
// converted into a lexical production with the magic prefix.
func nameReferenceForToken(tok *ebnf.Token) *ebnf.Name {
	if tok == nil {
		panic("Nil token provided. Tokens should never be nil.")
	}
	expectedName := literalNameFromLexeme(tok)
	replacement := &ebnf.Name{
		String: expectedName,
	}
	return replacement
}

func (gram *G) replaceAlternative(val ebnf.Alternative) {
	for i, alt := range val {
		if tok, ok := alt.(*ebnf.Token); ok {
			val[i] = nameReferenceForToken(tok)
			continue
		}
		gram.replaceExpressionIfLiteral(alt)
	}
}

func (gram *G) replaceSequence(val ebnf.Sequence) {
	for i, seq := range val {
		if tok, ok := seq.(*ebnf.Token); ok {
			val[i] = nameReferenceForToken(tok)
			continue
		}
		gram.replaceExpressionIfLiteral(seq)
	}
}

func (gram *G) replaceRepetition(val *ebnf.Repetition) {
	if tok, ok := val.Body.(*ebnf.Token); ok {
		val.Body = nameReferenceForToken(tok)
		// Skip the recursive step.
		// Since we already know what type the body is
		// and that it doesn't need to be recursed into.
		return
	}

	gram.replaceExpressionIfLiteral(val.Body)
}

func (gram *G) replaceOption(val *ebnf.Option) {
	if tok, ok := val.Body.(*ebnf.Token); ok {
		val.Body = nameReferenceForToken(tok)
		// Skip the recursive step.
		// Since we already know what type the body is
		// and that it doesn't need to be recursed into.
		return
	}

	gram.replaceExpressionIfLiteral(val.Body)
}

func (gram *G) replaceGroup(val *ebnf.Group) {
	if tok, ok := val.Body.(*ebnf.Token); ok {
		val.Body = nameReferenceForToken(tok)
		// Skip the recursive step.
		// Since we already know what type the body is
		// and that it doesn't need to be recursed into.
		return
	}
	gram.replaceExpressionIfLiteral(val.Body)
}

// expr is the expression whose composite parts we want to replace
// We want to find any token expression and replace it with an
// approperate production.
func (gram *G) replaceExpressionIfLiteral(expr ebnf.Expression) {
	switch val := expr.(type) {
	case ebnf.Alternative:
		gram.replaceAlternative(val)
	case ebnf.Sequence:
		gram.replaceSequence(val)
	case *ebnf.Repetition:
		gram.replaceRepetition(val)
	case *ebnf.Option:
		gram.replaceOption(val)
	case *ebnf.Group:
		gram.replaceGroup(val)
	}
}
