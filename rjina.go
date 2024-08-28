package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

  url := "https://r.jina.ai/"
	var userInput string
	
	fmt.Scanln(&userInput)
	resp, _ := http.Get(url + userInput)
	if resp.StatusCode == 200 {
		fmt.Println("Response OK")
		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))
	} else {
		fmt.Println("Response Failed")
	}
}
