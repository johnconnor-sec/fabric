package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func main() {
	var userInput string
	flag.StringVar(&userInput, "r", "", "Specify input for the URL")
	flag.Parse()

	url := "https://r.jina.ai/"
	resp, err := http.Get(url + userInput)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Response OK")
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("Response Failed")
	}
}

