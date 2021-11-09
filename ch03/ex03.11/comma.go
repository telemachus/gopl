package ch03

import (
	"bytes"
	"strings"
)

func Commify(s string) string {
	// Return empty string immediately.
	if len(s) == 0 {
		return s
	}

	start := 0
	end := len(s)
	var sign, number, fractionalPart string

	// Handle possible sign.
	if hasSign(s) {
		sign = string(s[0])
		start = 1
	}

	// Handle possible fractional part.
	if i := strings.Index(s, "."); i != -1 {
		fractionalPart = s[i:end]
		end = i
	}

	// Place commas in number.
	number = addCommas(s[start:end])

	// Return sign + number + fractional part.
	return sign + number + fractionalPart
}

func hasSign(s string) bool {
	return s[0] == '-' || s[0] == '+'
}

func addCommas(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	i := 0
	rem := n
	j := 1

	for i < n-1 {
		buf.WriteByte(s[i])
		rem = n - j
		if rem%3 == 0 {
			buf.WriteByte(',')
		}
		i++
		j++
	}
	buf.WriteByte(s[n-1])
	return buf.String()
}
