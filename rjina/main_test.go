package main

import (
	"io"
	"os"
	"net/http"
	"net/http/httptest"
	"testing"
)

var url string // Declare url variable globally

func main() {
	// Simulate main function logic that uses the url variable
	resp, err := http.Get(url)
	if err != nil {
		os.Stdout.Write([]byte("Response Error"))
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	os.Stdout.Write([]byte("Response OK\n" + string(body)))
}

func TestMain(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() == "/" { // Check for root path
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Test response"))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Override the default URL with the test server URL
	url = server.URL // Use the global url variable

	// Call the main function
	actualOutput := captureOutput(main)

	// Verify the output
	expectedOutput := "Response OK\nTest response" // Correct expected output format
	if actualOutput != expectedOutput {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, actualOutput)
	}
}

// Helper function to capture the output of the main function
func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

