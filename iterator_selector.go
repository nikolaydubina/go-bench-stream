package gostreambench

import (
	"bufio"
	"io"
)

type Iterator interface {
	Next() ([]byte, error)
}

type IteratorSelector struct {
	It   Iterator
	Vals map[string]bool
}

func (s IteratorSelector) Next() ([]byte, error) {
	v, err := s.It.Next()
	if err != nil {
		return nil, err
	}

	if len(v) == 0 {
		return nil, nil
	}

	if !s.Vals[string(v)] {
		return nil, nil
	}

	return v, nil
}

type IteratorSelectorFromReader struct {
	scanner *bufio.Scanner
}

func NewIteratorSelectorFromReader(r io.Reader) *IteratorSelectorFromReader {
	return &IteratorSelectorFromReader{
		scanner: bufio.NewScanner(r),
	}
}

func (s *IteratorSelectorFromReader) Next() ([]byte, error) {
	if s.scanner.Scan() {
		return s.scanner.Bytes(), nil
	}
	if err := s.scanner.Err(); err != nil {
		return nil, err
	}
	return nil, io.EOF
}

type ReaderFromIteratorSelector struct {
	it  Iterator
	buf []byte
}

func NewReaderFromIteratorSelector(it Iterator) *ReaderFromIteratorSelector {
	return &ReaderFromIteratorSelector{
		it:  it,
		buf: make([]byte, 0, 100),
	}
}

func (s *ReaderFromIteratorSelector) Read(b []byte) (n int, err error) {
	if len(s.buf) > 0 {
		l := copy(b, s.buf)
		s.buf = s.buf[l:]
		return l, nil
	}

	v, err := s.it.Next()
	if err != nil {
		return 0, err
	}
	if len(v) == 0 {
		return 0, nil
	}

	n = copy(b, v)

	if n < len(v) || (len(b) < (len(v) + 1)) {
		copy(s.buf, v[n:])
		s.buf = append(s.buf, '\n')
		return n, nil
	}

	b[n] = '\n'
	return n + 1, nil
}
