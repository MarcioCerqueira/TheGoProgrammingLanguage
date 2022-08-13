package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	b1 := []byte("x")
	b2 := []byte("X")
	if len(os.Args) > 1 {
		if os.Args[1] == "sha384" {
			c1 := sha512.Sum384(b1)
			c2 := sha512.Sum384(b2)
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
			return
		} else if os.Args[1] == "sha512" {
			c1 := sha512.Sum512(b1)
			c2 := sha512.Sum512(b2)
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
			return
		}
	}
	c1 := sha256.Sum256(b1)
	c2 := sha256.Sum256(b2)
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}
