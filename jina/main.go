package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func makeRequest(url string, inputStr string) (string, error) {
	fullURL := url + inputStr
	response, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Convert body to string
	bodyStr := string(body)

	// Use regex to remove markdown links
	linkRegex := regexp.MustCompile(`\!\[.*?\]\(.*?\)|\[.*?\]\(.*?\)`)
	bodyStr = linkRegex.ReplaceAllString(bodyStr, "")

  // Use regex to collapse more than three consecutive blank lines
	blankLineRegex := regexp.MustCompile(`(\n\s*\n\s*\n)(\s*\n)+`)
	bodyStr = blankLineRegex.ReplaceAllString(bodyStr, "")

   	specialCharRegex := regexp.MustCompile(`(?m)^\s*\*\s*$`)
	bodyStr = specialCharRegex.ReplaceAllString(bodyStr, "")

 	return bodyStr, nil

}

func main() {
	// Parse command-line arguments
	readURL := flag.String("r", "", "The URL to be sent to the API")
	searchInput := flag.String("s", "", "The input string to be sent to the API")
	flag.Parse()

	if *readURL == "" && *searchInput == "" {
		fmt.Println("Please provide either -r for read URL or -s for search input")
		return
	}

	var url, input string
	if *readURL != "" {
		url = "https://r.jina.ai/"
		input = *readURL
	} else {
		url = "https://s.jina.ai/"
		input = *searchInput
	}

	// Make the request
	response, err := makeRequest(url, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Print the response
	fmt.Println(response)
}
