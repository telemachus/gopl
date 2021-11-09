package echo

import "strings"

func Concat(ss []string) string {
	var s, sep string
	for _, arg := range ss {
		s += sep + arg
		sep = " "
	}
	return s
}

func Join(ss []string) string {
	return strings.Join(ss, " ")
}
