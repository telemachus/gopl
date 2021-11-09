package newreader

import "io"

type stringReader struct {
	s string
}

func (sr *stringReader) Read(p []byte) (int, error) {
	var err error
	var n int

	n = copy(p, sr.s)
	sr.s = sr.s[n:]
	if len(sr.s) == 0 {
		err = io.EOF
	}

	return n, err
}

func NewReader(s string) io.Reader {
	return &stringReader{s}
}
