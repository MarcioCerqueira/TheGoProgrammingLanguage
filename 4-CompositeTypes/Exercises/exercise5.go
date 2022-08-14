package main

import "fmt"

func removeDups(s []string) []string {
	j := 1
	for i, v := range s {
		if i > 0 && v != s[i-1] {
			s[j] = s[i]
			j++
		}
	}
	return s[:j]
}

func main() {
	s := []string{"Marcio", "Marcio", "Cerqueira", "de", "de"}
	fmt.Println(removeDups(s))
}
