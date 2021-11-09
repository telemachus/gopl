package ch04

import (
	"unicode"
	"unicode/utf8"
)

func SquashRune(s string) string {
	i := 0
	r := []rune(s)
	var last rune

	for _, c := range r {
		if !unicode.IsSpace(c) {
			r[i] = c
			i++
		} else if unicode.IsSpace(c) && !unicode.IsSpace(last) {
			r[i] = ' '
			i++
		}
		last = c
	}
	r = r[:i]

	return string(r)
}

func SquashByte(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}
