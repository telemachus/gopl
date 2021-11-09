package ch04_test

import (
	"ch04"
	"testing"
)

var phrase = "Hello, 世界!"
var reversePhrase = "!界世 ,olleH"

func TestReverseRune(t *testing.T) {
	expected := reversePhrase
	actual := string(ch04.ReverseRune([]rune(phrase)))

	if expected != actual {
		t.Errorf("expected %q; actual %q", expected, actual)
	}
}

func TestReverseByte(t *testing.T) {
	expected := reversePhrase
	actual := string(ch04.ReverseByte([]byte(phrase)))

	if expected != actual {
		t.Errorf("expected %q; actual %q", expected, actual)
	}
}

func BenchmarkReverseRune(b *testing.B) {
	r := []rune(phrase)
	for i := 0; i < b.N; i++ {
		ch04.ReverseRune(r)
	}
}

func BenchmarkReverseByte(b *testing.B) {
	bytes := []byte(phrase)
	for i := 0; i < b.N; i++ {
		ch04.ReverseByte(bytes)
	}
}
