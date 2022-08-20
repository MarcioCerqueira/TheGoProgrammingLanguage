package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

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
	doc, err := html.Parse(bytes.NewReader(fetch(os.Args[1])))
	if err != nil {
		fmt.Fprintf(os.Stderr, " %v\n", err)
		os.Exit(1)
	}
	outline(doc)
}

func outline(n *html.Node) {
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		fmt.Println(strings.TrimSpace(n.Data))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(c)
	}
}
