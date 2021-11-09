package anagram

import "strings"

func IsAnagram(s1, s2 string) bool {
	wordOne := strings.ToLower(s1)
	wordTwo := strings.ToLower(s2)
	if len(wordOne) != len(wordTwo) || wordOne == wordTwo {
		return false
	}

	chars := make(map[rune]int)
	for _, c := range wordOne {
		chars[c]++
	}
	for _, c := range wordTwo {
		chars[c]--
	}

	for _, count := range chars {
		if count != 0 {
			return false
		}
	}

	return true

}
