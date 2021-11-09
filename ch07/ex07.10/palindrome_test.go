package palindrome_test

import (
	"gopl/ch07/ex07.10/palindrome"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := map[string]struct {
		sequence string
		expected bool
	}{
		"an empty string is a palindrome":  {"", true},
		"most strings are not palindromes": {"not a palindrome", false},
		"“abba”":                           {"abba", true},
		"“Abba”":                           {"Abba", true},
		"“A man a plan a canal Panama”":    {"AmanaplanacanalPanama", true},
		"“1001 1001”":                      {"1001 1001", true},
		"“ 世界 界世 ”":                    {" 世界界世 ", true},
		"“Hello, 世界界世 ,olleh”":         {"Hello, 世界界世 ,olleh", true},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			w := palindrome.Word([]rune(tc.sequence))
			actual := palindrome.IsPalindrome(w)
			if tc.expected != actual {
				t.Errorf("expected %t; actual %t; string %q", tc.expected, actual, tc.sequence)
			}
		})
	}
}
