package gostreambench

import (
	"bufio"
	"io"
)

// ReaderSelector keeps lines that equal to string. Skip empty lines.
type ReaderSelector struct {
	scanner *bufio.Scanner
	buf     []byte
	vals    map[string]bool
}

func NewReaderSelector(r io.Reader, vals map[string]bool) *ReaderSelector {
	var l int
	for k := range vals {
		if len(k) > l {
			l = len(k)
		}
	}
	return &ReaderSelector{
		scanner: bufio.NewScanner(r),
		buf:     make([]byte, 0, l+1), // max len of selected val plus new line char
		vals:    vals,
	}
}

func (s *ReaderSelector) Read(b []byte) (n int, err error) {
	if len(s.buf) > 0 {
		l := copy(b, s.buf)
		s.buf = s.buf[l:]
		return l, nil
	}

	if s.scanner.Scan() {
		q := s.scanner.Bytes()

		if len(q) == 0 {
			return 0, nil
		}

		if !s.vals[string(q)] {
			return 0, nil
		}

		n = copy(b, q)

		if n < len(q) || (len(b) < (len(q) + 1)) {
			copy(s.buf, q[n:])
			s.buf = append(s.buf, '\n')
			return n, nil
		}

		b[n] = '\n'
		return n + 1, nil
	}

	if err := s.scanner.Err(); err != nil {
		return 0, err
	}

	return 0, io.EOF
}
