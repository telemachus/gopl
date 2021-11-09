package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const prefix string = "http://"

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		err = resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
	}
}
