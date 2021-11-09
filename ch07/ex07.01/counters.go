package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type ByteCounter int
type LineCounter int
type WordCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var n int
	r := bytes.NewReader(p)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n++
	}

	if err := scanner.Err(); err != nil {
		// Do not alter c if there is an error
		return 0, err
	}
	*c += LineCounter(n)
	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	var n int
	r := bytes.NewReader(p)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		n++
	}

	if err := scanner.Err(); err != nil {
		// Do not alter c if there is an error
		return 0, err
	}
	*c += WordCounter(n)
	return len(p), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: counter words|lines|bytes\n")
		os.Exit(1)
	}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ouch: %v\n", err)
		os.Exit(1)
	}
	choice := os.Args[1]
	switch choice {
	case "words":
		var wc WordCounter
		_, err := wc.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ouch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d words\n", wc)
	case "lines":
		var lc LineCounter
		_, err := lc.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ouch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d lines\n", lc)
	case "bytes":
		var bc ByteCounter
		_, err := bc.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ouch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "%d bytes\n", bc)
	default:
		fmt.Fprintf(os.Stderr, "choice not recognized: %q\n", choice)
		fmt.Fprintf(os.Stderr, "usage: counter words|lines|bytes\n")
		os.Exit(1)
	}
	os.Exit(0)
}
