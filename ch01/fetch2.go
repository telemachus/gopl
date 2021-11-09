package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if needsPrefix(url) {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			continue
		}

		defer resp.Body.Close()
		fmt.Printf("fetch2: %s (%s)\n", url, resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Printf("fetch2: problem copying %v—%s\n", url, err)
		}
	}
}

func needsPrefix(url string) bool {
	needs := true
	if strings.HasPrefix(url, "https://") {
		needs = false
	} else if strings.HasPrefix(url, "http://") {
		needs = false
	}
	return needs
}
