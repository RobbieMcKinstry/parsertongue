package main

const eof = rune(0)

func (glob LexemeGlob) peek() rune {
	return eof
}

func (glob LexemeGlob) next() rune {
	return eof
}

func (glob LexemeGlob) backup() {
}
