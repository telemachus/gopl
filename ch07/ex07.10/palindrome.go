package palindrome

import (
	"fmt"
	"unicode"
)

type Word []rune

func (w Word) Len() int {
	return len(w)
}

func (w Word) Less(i, j int) bool {
	a := fmt.Sprintf("%U", unicode.ToLower(w[i]))
	b := fmt.Sprintf("%U", unicode.ToLower(w[j]))

	return (a < b) || (b < a)
}

func (w Word) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func IsPalindrome(w Word) bool {
	for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
		if w.Less(i, j) {
			return false
		}
	}
	return true
}
