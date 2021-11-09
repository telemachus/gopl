package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var showFiles bool

func init() {
	const usage = "show where duplicates appear"
	flag.BoolVar(&showFiles, "files", false, usage)
	flag.BoolVar(&showFiles, "f", false, usage+" (shorthand)")
}

func main() {
	counts := make(map[string]int)
	fileNames := make(map[string]string)

	flag.Parse()
	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "dup: Please provide filenames.\n")
		os.Exit(1)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, arg, counts, fileNames)
			f.Close()
		}
	}

	if showFiles {
		for l, n := range counts {
			if n > 1 {
				fmt.Printf("[%s] in %s: %d times\n", l, fileNames[l],n)
			}
		}
	} else {
		for l, n := range counts {
			if n > 1 {
				fmt.Printf("[%s] appears %d times\n", l, n)
			}
		}
	}
}

func countLines(f *os.File, arg string, counts map[string]int, fileNames map[string]string) {
	seenForArg := make(map[string]int)
	input := bufio.NewScanner(f)
	sep := ", "
	for input.Scan() {
		l := input.Text()
		counts[l]++
		seenForArg[l]++
		if seenForArg[l] < 2 {
			if fileNames[l] == "" {
				fileNames[l] = arg
			} else {
				fileNames[l] = fileNames[l] + sep + arg
			}
		}
	}
}
