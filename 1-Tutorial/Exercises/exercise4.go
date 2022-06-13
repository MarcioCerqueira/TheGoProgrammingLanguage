// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string]*os.File)
	printedFilenames := make(map[string]bool)
	filenames := os.Args[1:]
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise4: %v\n", err)
			continue
		}
		files[filename] = f
		printedFilenames[filename] = false
		countLines(f, counts)
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			for filename, _ := range files {
				if !printedFilenames[filename] {
					printedFilenames[filename] = printDupFilenames(filename, line)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	//NOTE: ignoring potential errors from input.Err()
}

func printDupFilenames(filename string, dupLine string) bool {
	f, _ := os.Open(filename)
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == dupLine {
			fmt.Println(f.Name())
			return true
		}
	}
	return false
}
