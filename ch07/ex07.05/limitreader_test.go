package limitreader_test

import (
	"gopl/ch07/ex07.05/limitreader"
	"io"
	"strings"
	"testing"
)

func TestLimitBelowZero(t *testing.T) {
	limit := int64(-5)
	s := strings.NewReader("Hello, LimitReaderWorld!")
	r := limitreader.LimitReader(s, limit)
	p := make([]byte, 0)
	n, err := r.Read(p)

	if io.EOF != err || n != 0 {
		t.Errorf("expected io.EOF and 0; actual %v and %d", err, n)
	}
}

func TestStringLongerThanLimit(t *testing.T) {
	limit := int64(4)
	s := "Hello, LimitReaderWorld!"
	r := limitreader.LimitReader(strings.NewReader(s), limit)
	p := make([]byte, limit)

	_, err := r.Read(p)
	if s[:limit] != string(p) {
		t.Errorf("expected %s; actual %s", s[:limit], string(p))
	}
	if nil != err {
		t.Errorf("expected nil; actual %v", err)
	}

	n, err := r.Read(p)
	if 0 != n || io.EOF != err {
		t.Errorf("expected 0 and io.EOF; actual %d and %v", n, err)
	}
}

func TestStringShorterThanLimit(t *testing.T) {
	limit := int64(8)
	s := "Hello"
	r := limitreader.LimitReader(strings.NewReader(s), limit)
	p := make([]byte, limit)

	n, err := r.Read(p)
	if s[:n] != string(p[:n]) {
		t.Errorf("expected %q; actual %q", s[:n], string(p[:n]))
	}
	if nil != err {
		t.Errorf("expected nil; actual %v", err)
	}

	n, err = r.Read(p)
	if 0 != n || io.EOF != err {
		t.Errorf("expected 0 and io.EOF; actual %d and %v", n, err)
	}
}
