package newreader_test

import (
	"bytes"
	"gopl/ch07/ex07.04/newreader"
	"io"
	"testing"
)

func TestNewReader(t *testing.T) {
	b := &bytes.Buffer{}
	s := "some io.reader stream to be read\n"
	r := newreader.NewReader(s)
	n, _ := io.Copy(b, r)
	if n != int64(b.Len()) {
		t.Errorf("read %d bytes; should have read %d bytes", n, b.Len())
	}

	if b.String() != s[:n] {
		t.Errorf("%s should be the same as %s", b.String(), s[:n])
	}
}
