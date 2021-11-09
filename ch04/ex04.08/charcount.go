package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	countsByType := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		switch {
		case r == unicode.ReplacementChar:
			invalid++
			continue
		case unicode.IsDigit(r):
			countsByType["digit"]++
		case unicode.IsLetter(r):
			countsByType["letter"]++
		case unicode.IsPunct(r):
			countsByType["punctuation"]++
		case unicode.IsSpace(r):
			countsByType["whitespace"]++
		}
	}

	fmt.Println("Characters by type:")
	for t, n := range countsByType {
		fmt.Printf("%25s\t%d\n", t, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
