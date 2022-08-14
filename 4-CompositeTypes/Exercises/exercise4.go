package main

import "fmt"

func rotate(s []int, shift int) {
	t := make([]int, len(s))
	for i, j := 0, shift; i < len(s); i, j = i+1, j+1 {
		if j >= len(s) {
			j = 0
		}
		t[i] = s[j]
	}
	copy(s, t)
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotate(a, 2)
	fmt.Println(a)
}
