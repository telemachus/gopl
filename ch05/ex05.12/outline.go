package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

const (
	progName    = "outline"
	exitSuccess = 0
	exitFailure = 1
)

func main() {
	exitValue := exitSuccess
	for _, url := range os.Args[1:] {
		if err := outline(url); err != nil {
			exitValue = exitFailure
			fmt.Fprintf(os.Stderr, "%s: %v\n", progName, err)
		}
	}

	os.Exit(exitValue)
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth, "", n.Data)
			depth += 4
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth -= 4
		}
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
