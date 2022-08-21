package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		fmt.Printf("Total word count: %d\n", words)
		fmt.Printf("Total images count: %d\n", images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		err = fmt.Errorf("getting %s: %s", url, resp.Status)
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

// visit appends to links each link found in n and returns the result.
func countWordsAndImages(n *html.Node) (words, images int) {
	var currentWordsCount, currentImageCount int
	if n.Type == html.ElementNode && n.Data == "img" {
		currentImageCount++
	}
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		currentWordsCount += len(strings.Split(strings.TrimSpace(n.Data), " "))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		childWordCount, childImageCount := countWordsAndImages(c)
		currentWordsCount += childWordCount
		currentImageCount += childImageCount
	}
	return currentWordsCount, currentImageCount
}
