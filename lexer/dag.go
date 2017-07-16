package lexer

import (
	"fmt"

	"golang.org/x/exp/ebnf"
)

// Convert the expression to a stateFn.
func (lex *L) toStateFn(exp ebnf.Expression) StateFn {
	switch v := exp.(type) {
	case *ebnf.Option:
		return lex.makeOption(v)
	case *ebnf.Group:
		return lex.makeGroup(v)
	case *ebnf.Repetition:
		return lex.makeRepetition(v)
	case ebnf.Sequence:
		return lex.makeSequence(v)
	case ebnf.Alternative:
		return lex.makeAlternative(v)
	case *ebnf.Name:
		return lex.makeName(v)
	case *ebnf.Token:
		return lex.makeToken(v)
	default:
		return nil
	}
}

func (lex *L) makeName(name *ebnf.Name) StateFn {

	// look up the name of the production
	if runeMatcher, ok := prebuilt[name.String]; ok {
		return lex.makeRuneMatcher(runeMatcher)
	}

	// else, we have to generate the production from
	// it's children according to the grammar.
	prod := lex.gram.Prod(name.String)
	return lex.toStateFn(prod.Expr)
}

// match on a single rune
func (lex *L) makeRuneMatcher(matcher runeMatcher) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		if matcher(lex.next()) {
			return nil, 1
		}
		lex.backup()
		return nil, 0
	}
}

// Logical AND
func (lex *L) makeSequence(seq ebnf.Sequence) StateFn {

	matchers := []StateFn{}
	for _, exp := range seq {
		matchers = append(matchers, lex.toStateFn(exp))
	}

	return func(lex *L, start int) (StateFn, int) {
		fmt.Printf("Running sequence %v\n", seq)
		var size = 0
		for i, match := range matchers {
			next := match.Exhaust(lex.Clone(), start)
			if next == 0 && !isOptional(seq[i]) {
				return nil, 0
			}
			size += next
			lex.advance(next)
		}
		return nil, size
	}
}

// Logical OR
func (lex *L) makeAlternative(alt ebnf.Alternative) StateFn {

	matchers := []StateFn{}
	for _, exp := range alt {
		matchers = append(matchers, lex.toStateFn(exp))
	}

	// TODO this can be converted to a parallel implementation
	return func(lex *L, start int) (StateFn, int) {

		var max = 0
		for _, match := range matchers {
			width := match.Exhaust(lex.Clone(), start)
			if width > max {
				max = width
			}
		}

		return nil, max
	}
}

func (lex *L) makeToken(tok *ebnf.Token) StateFn {
	// for a token, read in the runes and check to make sure they match.
	literal := tok.String

	return func(lex *L, start int) (StateFn, int) {

		for _, char := range literal {
			nextRune := lex.next()
			if nextRune != char {
				return nil, 0
			}
		}

		return nil, len(literal)
	}
}

func (lex *L) makeGroup(group *ebnf.Group) StateFn {
	return lex.toStateFn(group.Body)
}

func (lex *L) makeOption(op *ebnf.Option) StateFn {
	var exp = op.Body
	var matcher = lex.toStateFn(exp)

	next := func(lex *L, start int) (StateFn, int) {
		match := matcher.Exhaust(lex, start)
		return nil, match
	}
	return next
}

func (lex *L) makeRepetition(rep *ebnf.Repetition) StateFn {
	var (
		exp     = rep.Body
		matcher = lex.toStateFn(exp)
	)

	var next = func(lex *L, start int) (StateFn, int) {
		var total = 0

		for size := matcher.Exhaust(lex.Clone(), start); size > 0; size = matcher.Exhaust(lex.Clone(), start+size) {

			total += size
			lex.advance(size)
		}

		return nil, total
	}
	return next
}

// Exhaust runs the statefn until it is nil, returning the
// final result
func (state StateFn) Exhaust(lex *L, start int) int {
	pos := start
	fn := state

	for fn != nil {
		fn, pos = fn(lex, pos)
	}
	return pos
}

func isOptional(exp ebnf.Expression) bool {
	_, optional := exp.(*ebnf.Option)
	_, repetition := exp.(*ebnf.Repetition)

	return optional || repetition
}
