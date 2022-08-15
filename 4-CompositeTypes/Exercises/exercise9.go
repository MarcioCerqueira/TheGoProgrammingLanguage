package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	wordfreq := make(map[string]int)
	f, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		wordfreq[scanner.Text()]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for key, value := range wordfreq {
		fmt.Printf("%s: %d\n", key, value)
	}
}
