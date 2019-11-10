package main

import (
	"io"
)

type ByteCounter int

func (b *ByteCounter) Write(bytes []byte) (int, error) {
	*b += ByteCounter(len(bytes)) // 转换类型
	return len(bytes), nil
}

type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}
type Closer interface {
	Close() error
}

// type ReadWriter interface {
// 	Reader
// 	Writer
// }

type ReadWriter interface {
	Reader
	Write([]byte) (int, error)
}

type ReadWriteCloser interface {
	ReadWriter
	Closer
}

type StringReader struct {
	str []byte
	p   int
}

func (s *StringReader) Read(b []byte) (int, error) {
	start := s.p
	if start >= len(s.str) {
		return 0, io.EOF
	}
	end := start + len(b)
	if end > len(s.str) {
		end = len(s.str)
	}
	copy(b, s.str[start:end])
	s.p += len(b)
	return len(b), nil
}

func NewStringReader(s string) *StringReader {
	return &StringReader{str: []byte(s)}
}
