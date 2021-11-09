package ch04_test

import (
	"ch04"
	"testing"
)

func isSameSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestFilterDups(t *testing.T) {
	tests := map[string]struct {
		strings  []string
		expected []string
	}{
		"empty slice": {strings: []string{}, expected: []string{}},
		"no dups":     {strings: []string{"foo", "bar"}, expected: []string{"foo", "bar"}},
		"one dup":     {strings: []string{"foo", "foo"}, expected: []string{"foo"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ch04.FilterDups(tc.strings)
			if !isSameSlice(tc.expected, actual) {
				t.Errorf("expected %#v; actual %#v", tc.expected, actual)
			}
		})
	}
}
