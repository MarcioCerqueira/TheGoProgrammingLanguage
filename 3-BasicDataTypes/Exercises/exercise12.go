package main

import (
	"fmt"
	"os"
)

func areAnagrams(firstString, secondString string) string {
	firstStringMap := make(map[rune]int)
	secondStringMap := make(map[rune]int)

	if len(firstString) != len(secondString) {
		return "An anagram was not found"
	}

	for _, value := range firstString {
		firstStringMap[value]++
	}

	for _, value := range secondString {
		secondStringMap[value]++
	}

	for key := range firstStringMap {
		if firstStringMap[key] != secondStringMap[key] {
			return "An anagram was not found"
		}
	}

	return "An anagram was found"

}

func main() {
	fmt.Printf(areAnagrams(os.Args[1], os.Args[2]))
}
