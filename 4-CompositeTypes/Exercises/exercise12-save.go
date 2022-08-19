package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Transcript struct {
	Transcript string `json:"transcript"`
}

type URLToTranscript struct {
	URL        string `json:"url"`
	Transcript string `json:"transcript"`
}

func main() {
	var URLs []URLToTranscript
	for i := 1; i <= 400; i++ {
		var CurrentURLToTranscript URLToTranscript
		CurrentURLToTranscript.URL = "http://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(CurrentURLToTranscript.URL)
		if err != nil {
			log.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			_ = fmt.Errorf("Http request failed: %s", resp.Status)
		}
		var CurrentTranscript Transcript
		if err := json.NewDecoder(resp.Body).Decode(&CurrentTranscript); err != nil {
			resp.Body.Close()
			log.Fatal(err)
		}
		CurrentURLToTranscript.Transcript = CurrentTranscript.Transcript
		URLs = append(URLs, CurrentURLToTranscript)
		resp.Body.Close()
		fmt.Println(CurrentURLToTranscript.URL)
	}
	data, err := json.Marshal(URLs)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("xkcd.json", data, os.ModePerm)
}
