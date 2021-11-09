package anagram_test

import (
	"anagram"
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := map[string]struct {
		s1       string
		s2       string
		expected bool
	}{
		"empty string": {s1: "", s2: "", expected: false},
		"identical words are not anagrams": {s1: "foo", s2: "foo", expected: false},
		"not anagrams unequal length": {s1: "foo", s2: "buzz", expected: false},
		"not anagrams equal length": {s1: "foo", s2: "bar", expected: false},
		"anagrams": {s1: "foo", s2: "oof", expected: true},
		"anagrams with accents": {s1: "fōö", s2: "ōöf", expected: true},
		"only differences in case = not anagram": {s1: "Tab", s2: "tab", expected: false},
		"differences in case and order = anagram": {s1: "foo", s2: "Oof", expected: true},
		"not anagrams second word longer": {s1: "tab", s2: "tabb", expected: false},
		"not anagrams first word longer": {s1: "tabb", s2: "tab", expected: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := anagram.IsAnagram(tc.s1, tc.s2)
			if tc.expected != actual {
				t.Errorf("expected %t; actual %t", tc.expected, actual)
			}
		})
	}
}
