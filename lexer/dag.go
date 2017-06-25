package lexer

import "golang.org/x/exp/ebnf"

// TODO
// Convert the expression to a stateFn.
func toStateFn(exp ebnf.Expression) StateFn {
	switch v := exp.(type) {
	case *ebnf.Option:
		return makeOption(v)
	case *ebnf.Group:
		return makeGroup(v)
	case *ebnf.Repetition:
		return makeRepetition(v)
	case ebnf.Sequence:
		return makeSequence(v)
	case ebnf.Alternative:
		return makeAlternative(v)
	case *ebnf.Name:
		return makeName(v)
	case *ebnf.Token:
		return makeToken(v)
	default:
		return nil

	}
}

// TODO impl
func makeName(name *ebnf.Name) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		return nil, 0
	}
}

// TODO impl
func makeSequence(seq ebnf.Sequence) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		return nil, 0
	}
}

// TODO impl
func makeAlternative(alt ebnf.Alternative) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		return nil, 0
	}
}

// TODO impl
func makeToken(tok *ebnf.Token) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		return nil, 0
	}
}

// TODO impl
func makeGroup(group *ebnf.Group) StateFn {
	return func(lex *L, start int) (StateFn, int) {
		return nil, 0
	}
}

func makeOption(op *ebnf.Option) StateFn {
	var (
		state   StateFn
		match   int
		exp     = op.Body
		matcher = toStateFn(exp)
	)

	next := func(lex *L, start int) (StateFn, int) {
		state, match = matcher, start
		for state != nil {
			state, match = state(lex, start)
		}
		return nil, match
	}
	return next
}

func makeRepetition(rep *ebnf.Repetition) StateFn {
	var (
		exp     = rep.Body
		matcher = toStateFn(exp)
	)

	var next = func(lex *L, start int) (StateFn, int) {
		size := matcher.Exhaust(lex, start)
		for i := size; i < 0; i = matcher.Exhaust(lex, i) {
			size += i
		}
		// TODO do I return nil here? I'm not sure…
		return nil, size
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
