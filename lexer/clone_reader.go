package lexer

import (
	"bytes"
	"io"
)

// A CloneScanner is a io.RuneScanner that is easy to clone
type CloneScanner interface {
	io.RuneScanner
	Clone() CloneScanner
}

// A BufferScanner is a CloneScanner which efficiently
// clones and allows concurrent reading.
type BufferScanner struct {
	*bytes.Reader
	data []byte
}

// NewBufferScanner news up a buffer.
func NewBufferScanner(payload []byte) *BufferScanner {
	scanner := new(BufferScanner)
	scanner.Reader = bytes.NewReader(payload)
	scanner.data = payload
	return scanner
}

// Clone returns a concurrently readable BufferScanner with
// neither having been mutated from the cloning
func (scanner *BufferScanner) Clone() CloneScanner {
	size := scanner.Size()
	length := int64(scanner.Len())
	pos := size - length

	res := NewBufferScanner(scanner.data)
	res.Reader.Seek(pos, io.SeekStart)
	return res
}

var _ CloneScanner = &BufferScanner{}
