package grammar

import (
	"fmt"
	"strings"

	"golang.org/x/exp/ebnf"
)

// clean iterates through the grammar, looking for all
// token literals. It replaces all token literals with
// productions which point to a single token. It adds
// those productions to the grammar.
// The produced productions are all lexical, since all
// token literals are entrant by definition.
func (gram *G) clean() {
	// Collect the tokens
	tokenMap := gram.FindTokenLiterals()
	// Now, for each of those tokens,
	// create a new production that contains only that token literal

	var productions = map[string]*ebnf.Production{}
	for _, tok := range tokenMap {
		tokenProduction := marshalIntoProduction(tok)
		// collect it for later
		productions[tokenProduction.Name.String] = tokenProduction
	}

	gram.replaceTokenLiterals(tokenMap)

	// finally, add each of the productions to the grammar:
	for name, prod := range productions {
		// check to see of the existing production is already
		// at the top level.
		gram.gram[name] = prod
	}
}

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

func literalNameFromLexeme(tok *ebnf.Token) string {
	body := tok.String
	edited := fmt.Sprintf("\u200B_lexeme_%v", body)
	return edited
}

func lexemeNameFromLiteral(tok *ebnf.Token) string {
	body := tok.String
	// Chop off the first couple of characters.
	edited := strings.Replace(body, "\u200B_lexeme_", "", 1)
	return edited
}

func (gram *G) replaceTokenLiterals(
	oldToks map[string]*ebnf.Token,
) {
	for _, production := range gram.gram {
		expr := production.Expr
		if _, ok := expr.(*ebnf.Token); ok {
			fmt.Println("Successfully found a top level production!")
			fmt.Println(Stringify(production))
			continue
		}
		gram.replaceExpressionIfLiteral(expr, oldToks)
	}
}

func (gram *G) replaceExpressionIfLiteral(
	expr ebnf.Expression,
	oldToks map[string]*ebnf.Token,
) {
	switch val := expr.(type) {
	case ebnf.Alternative:
		for i, alt := range val {
			if tok, ok := alt.(*ebnf.Token); ok {
				tokLit := oldToks[tok.String]
				expectedName := literalNameFromLexeme(tokLit)
				replacement := &ebnf.Name{
					String: expectedName,
				}

				val[i] = replacement
				continue
			}
			gram.replaceExpressionIfLiteral(alt, oldToks)
		}
	case ebnf.Sequence:
		for i, seq := range val {
			if tok, ok := seq.(*ebnf.Token); ok {
				tokLit := oldToks[tok.String]
				expectedName := literalNameFromLexeme(tokLit)
				replacement := &ebnf.Name{
					String: expectedName,
				}
				val[i] = replacement
				continue
			}
			gram.replaceExpressionIfLiteral(seq, oldToks)
		}
	case *ebnf.Repetition:
		if tok, ok := val.Body.(*ebnf.Token); ok {
			tokLit := oldToks[tok.String]
			expectedName := literalNameFromLexeme(tokLit)
			replacement := &ebnf.Name{
				String: expectedName,
			}
			val.Body = replacement
		} else {
			gram.replaceExpressionIfLiteral(val.Body, oldToks)
		}
	case *ebnf.Option:
		if tok, ok := val.Body.(*ebnf.Token); ok {
			tokLit := oldToks[tok.String]
			expectedName := literalNameFromLexeme(tokLit)
			replacement := &ebnf.Name{
				String: expectedName,
			}
			val.Body = replacement
		} else {
			gram.replaceExpressionIfLiteral(val.Body, oldToks)
		}
	case *ebnf.Group:
		if tok, ok := val.Body.(*ebnf.Token); ok {
			tokLit := oldToks[tok.String]
			expectedName := literalNameFromLexeme(tokLit)
			replacement := &ebnf.Name{
				String: expectedName,
			}
			val.Body = replacement
		} else {
			gram.replaceExpressionIfLiteral(val.Body, oldToks)
		}
	}
}
