package main

import (
	"strings"
	"unicode/utf8"
)

const eof = rune(0)

func (glob LexemeGlob) peek() rune {
	r := glob.next()
	glob.backup()
	return r
}

func (glob LexemeGlob) next() rune {
	if glob.end >= len(*glob.file) {
		glob.width = 0
		return eof
	}
	var r rune
	var str string = *glob.file
	r, glob.width = utf8.DecodeRuneInString(str[glob.end:])
	glob.end += glob.width
	return r
}

func (glob LexemeGlob) backup() {
	glob.end -= glob.width
}

func (glob LexemeGlob) ignore() {
	glob.start = glob.end
}

func (glob LexemeGlob) size() int {
	return glob.end - glob.start
}

// accept looks to see if the next rune is contained in the valid string
func (glob LexemeGlob) accept(valid string) bool {
	for strings.IndexRune(valid, glob.next()) >= 0 {
		return true
	}
	glob.backup()
	return false
}

func (glob LexemeGlob) acceptRune(valid string) {
	for strings.IndexRune(valid, glob.next()) >= 0 {
	}
	glob.backup()
}
