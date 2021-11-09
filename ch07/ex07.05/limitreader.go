package limitreader

import "io"

type limitReader struct {
	r io.Reader
	n int64
}

func (lr *limitReader) Read(p []byte) (int, error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}

	if int64(len(p)) > lr.n {
		p = p[:lr.n]
	}
	n, err := lr.r.Read(p)
	lr.n -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r: r, n: n}
}
