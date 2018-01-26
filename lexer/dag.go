package lexer

import (
	"fmt"

	"github.com/RobbieMcKinstry/parsertongue/grammar"

	"golang.org/x/exp/ebnf"
)

// Convert the expression to a stateFn.
func (lex *L) toStateFn(exp ebnf.Expression) StateFn {

	fmt.Println("Building state func:")
	fmt.Println(grammar.ExprString(exp))
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
		panic("No such implementation")
	}
}

func stringifyExpr(exp ebnf.Expression) string {
	switch v := exp.(type) {
	case *ebnf.Option:
		return fmt.Sprintf("[%s]", stringifyExpr(v.Body))
	case *ebnf.Group:
		return fmt.Sprintf("(%s)", stringifyExpr(v.Body))
	case *ebnf.Repetition:
		return fmt.Sprintf("[%s]", stringifyExpr(v.Body))
	case ebnf.Sequence:
		strings := make([]string, 0, len(v))
		// collect the stringified subexpressions
		for _, exp := range v {
			strings = append(strings, stringifyExpr(exp))
		}
		res := strings[0]
		for _, str := range strings[1:] {
			res += fmt.Sprintf(" + %s", str)
		}
		return res
	case ebnf.Alternative:
		alts := make([]string, 0, len(v))
		// collect the stringified alternatives
		for _, exp := range v {
			alts = append(alts, stringifyExpr(exp))
		}
		res := alts[0]
		for _, str := range alts[1:] {
			res += fmt.Sprintf(" | %s", str)
		}
		return res
	case *ebnf.Name:
		return v.String
	case *ebnf.Token:
		return fmt.Sprintf("\"%s\"", v.String)
	default:
		return "error"
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
	return func(lex *L) (StateFn, int) {
		if matcher(lex.next()) {
			return nil, 1
		}
		lex.backup()
		return nil, -1
	}
}

// Logical AND
func (lex *L) makeSequence(seq ebnf.Sequence) StateFn {

	matchers := []StateFn{}
	for _, exp := range seq {
		matchers = append(matchers, lex.toStateFn(exp))
	}

	return func(lex *L) (StateFn, int) {
		var size = 0
		for i, match := range matchers {
			next := match.Exhaust(lex.Clone())
			if next == -1 && !isOptional(seq[i]) {
				return nil, -1
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
	return func(lex *L) (StateFn, int) {

		var max = -1
		for _, match := range matchers {
			width := match.Exhaust(lex.Clone())

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

	return func(lex *L) (StateFn, int) {

		for _, char := range literal {
			nextRune := lex.next()
			if nextRune != char {
				return nil, -1
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

	next := func(lex *L) (StateFn, int) {
		match := matcher.Exhaust(lex)
		if match == -1 {
			return nil, 0
		}
		return nil, match
	}
	return next
}

func (lex *L) makeRepetition(rep *ebnf.Repetition) StateFn {
	var (
		exp     = rep.Body
		matcher = lex.toStateFn(exp)
	)

	var next = func(lex *L) (StateFn, int) {
		var total = 0

		for size := matcher.Exhaust(lex.Clone()); size > 0; size = matcher.Exhaust(lex.Clone()) {

			total += size
			lex.advance(size)
		}

		return nil, total
	}
	return next
}

// Exhaust runs the statefn until it is nil, returning the
// final result
func (state StateFn) Exhaust(lex *L) int {
	size := 0
	for pos, fn := 0, state; fn != nil; size += pos {
		fn, pos = fn(lex)
		if pos == -1 {
			return -1
		}
	}
	return size
}

func isOptional(exp ebnf.Expression) bool {
	_, optional := exp.(*ebnf.Option)
	_, repetition := exp.(*ebnf.Repetition)

	return optional || repetition
}
