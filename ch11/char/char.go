// Package char provides utility functions for Unicode characters.
package char

import "unicode"

// Count returns a map counting Unicode characters in a string and an integer
// with the number of invalid characters in the string.
func Count(s string) (map[rune]int, int) {
	counts := make(map[rune]int)
	invalid := 0
	for _, r := range s {
		if r == unicode.ReplacementChar {
			invalid++
		} else {
			counts[r]++
		}
	}
	return counts, invalid
}
