package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type URLToTranscript struct {
	URL        string `json:"url"`
	Transcript string `json:"transcript"`
}

func main() {
	file, err := ioutil.ReadFile("xkcd.json")
	if err != nil {
		log.Fatal(err)
	}
	var data []URLToTranscript
	err = json.Unmarshal([]byte(file), &data)
	for _, value := range data {
		if strings.Contains(value.Transcript, os.Args[1]) {
			fmt.Println(value)
		}
	}
}
