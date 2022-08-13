package main

import (
	"crypto/sha256"
	"fmt"
)

//pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func CountBitDifference(c1, c2 [32]byte) int {
	var count int
	for b := 0; b < len(c1); b++ {
		count += int(pc[c1[b]^c2[b]])
	}
	return count
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%d\n", c1, c2, CountBitDifference(c1, c2))
}
