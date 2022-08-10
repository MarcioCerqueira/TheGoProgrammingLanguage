package main

import (
	"bytes"
	"fmt"
	"os"
)

//comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	firstComma := n % 3
	for i := 0; i < firstComma; i++ {
		buf.WriteByte(s[i])
	}
	if firstComma > 0 {
		buf.WriteByte(',')
	}
	count := 0
	for i := firstComma; i < n; i++ {
		if count%3 == 0 && count > 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
		count++
	}
	return buf.String()
}

func main() {
	fmt.Println(comma(os.Args[1]))
}
