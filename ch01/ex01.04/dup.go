package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts)
	} else {
		for _, fileName := range files {
			fh, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(fileName, fh, counts)
			fh.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s: %d\n", line, n)
		}
	}
}

func countLines(fileName string, fh *os.File, counts map[string]int) {
	input := bufio.NewScanner(fh)
	for input.Scan() {
		nameAndLine := fmt.Sprintf("%s: %q", fileName, input.Text())
		counts[nameAndLine]++
	}
}
