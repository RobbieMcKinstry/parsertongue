package lexer

import "testing"

func TestClone(t *testing.T) {
	var payload = []byte("hello world")

	r1 := NewBufferScanner(payload)

	// read the substring "hello "
	for i := 0; i < 6; i++ {
		r1.ReadRune()
	}
	// clone the reader, then try to read the rest of the line.
	r2 := r1.Clone()

	for range "world" {
		expected, _, _ := r1.ReadRune()
		observed, _, _ := r2.ReadRune()
		if expected != observed {
			t.Errorf("Expected %c but observed %c", expected, observed)
		}
	}
}
