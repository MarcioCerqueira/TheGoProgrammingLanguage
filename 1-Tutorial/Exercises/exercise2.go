package main

import (
	"fmt"
	"os"
)

func main() {
	for index, argument := range os.Args[0:] {
		fmt.Println(index, argument)
	}
}
