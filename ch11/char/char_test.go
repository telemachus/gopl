package char_test

import (
	"gopl/ch11/char"
	"testing"
)

func TestCountInvalid(t *testing.T) {
	expected := 3
	s := "a�a�a�"
	_, actual := char.Count(s)

	if expected != actual {
		t.Errorf("input string: %q, invalid = %d", s, actual)
	}
}

func TestCountValid(t *testing.T) {
	s := "aaa"
	expected := make(map[rune]int)
	expected['a'] = 3
	actual, _ := char.Count(s)

	for r, n := range expected {
		if actual[r] != n {
			t.Errorf("input string: %q, but result['%c'] != %d", s, r, n)
		}
	}
}
