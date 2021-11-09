package ch03

import (
	"bytes"
)

func Comma(s string) string {
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
