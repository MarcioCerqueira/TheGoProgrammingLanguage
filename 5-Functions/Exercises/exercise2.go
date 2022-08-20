package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func fetch(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	return b
}

func main() {
	elementNames := make(map[string]int)
	doc, err := html.Parse(bytes.NewReader(fetch(os.Args[1])))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	outline(elementNames, doc)
	for key, value := range elementNames {
		fmt.Println(key, value)
	}
}

func outline(elementNames map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		elementNames[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(elementNames, c)
	}
}
