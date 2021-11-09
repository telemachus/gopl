package ch04_test

import (
	"ch04"
	"testing"
)

var benchString = "\t\n   foobar  foo\tbar\\fizz   \nbuzz\n\t  "
var benchBytes = []byte(benchString)

func TestSquashRune(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected string
	}{
		"empty string": {s: "", expected: ""},
		"two spaces": {s: "foo  bar", expected: "foo bar"},
		"tab": {s: "foo\tbar", expected: "foo bar"},
		"two tabs": {s: "foo\t\tbar", expected: "foo bar"},
		"single space at start of string": {s: " foobar", expected: " foobar"},
		"two spaces at start of string": {s: "  foobar", expected: " foobar"},
		"single space at end of string": {s: "foobar ", expected: "foobar "},
		"two spaces at end of string": {s: "foobar  ", expected: "foobar "},
		"newline mid-string": {s: "foo\nbar", expected: "foo bar"},
		"newline at start of string": {s: "\nfoobar", expected: " foobar"},
		"newline at end of string": {s: "foobar\n", expected: "foobar "},
		"a mix of everything everywhere": {s: "\t\n   foobar  foo\tbar\\fizz   \nbuzz\n\t  ", expected: " foobar foo bar\\fizz buzz "},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ch04.SquashRune(tc.s)
			if tc.expected != actual {
				t.Errorf("expected %q; actual %q", tc.expected, actual)
			}
		})
	}
}

func TestSquashByte(t *testing.T) {
	tests := map[string]struct {
		s        string
		expected string
	}{
		"empty string": {s: "", expected: ""},
		"two spaces": {s: "foo  bar", expected: "foo bar"},
		"tab": {s: "foo\tbar", expected: "foo bar"},
		"two tabs": {s: "foo\t\tbar", expected: "foo bar"},
		"single space at start of string": {s: " foobar", expected: " foobar"},
		"two spaces at start of string": {s: "  foobar", expected: " foobar"},
		"single space at end of string": {s: "foobar ", expected: "foobar "},
		"two spaces at end of string": {s: "foobar  ", expected: "foobar "},
		"newline mid-string": {s: "foo\nbar", expected: "foo bar"},
		"newline at start of string": {s: "\nfoobar", expected: " foobar"},
		"newline at end of string": {s: "foobar\n", expected: "foobar "},
		"a mix of everything everywhere": {s: "\t\n   foobar  foo\tbar\\fizz   \nbuzz\n\t  ", expected: " foobar foo bar\\fizz buzz "},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			b := []byte(tc.s)
			actual := string(ch04.SquashByte(b))
			if tc.expected != actual {
				t.Errorf("expected %q; actual %q", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSquashRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch04.SquashRune(benchString)
	}
}

func BenchmarkSquashByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch04.SquashByte(benchBytes)
	}
}
