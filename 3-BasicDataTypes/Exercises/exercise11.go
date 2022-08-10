package main

import (
	"bytes"
	"fmt"
	"os"
)

//comma inserts commas in a decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	sign := 0
	n := len(s)
	if n > 0 && (s[0] == '+' || s[0] == '-') {
		sign = 1
		buf.WriteByte(s[0])
	}
	if (n - sign) <= 3 {
		return s
	}
	firstComma := (n - sign) % 3
	for i := sign; i < firstComma+sign; i++ {
		buf.WriteByte(s[i])
	}
	if firstComma > 0 {
		buf.WriteByte(',')
	}
	count := 0
	for i := firstComma + sign; i < n; i++ {
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
