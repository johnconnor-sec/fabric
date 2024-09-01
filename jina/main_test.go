package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestMakeRequest tests the makeRequest function
func TestMakeRequest(t *testing.T) {
    // Create a test server
    ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, client"))
    }))
    defer ts.Close()

    // Test cases
    tests := []struct {
        url      string
        inputStr string
        expected string
    }{
        {ts.URL, "", "Hello, client"},
        {ts.URL, "/test", "Hello, client"},
    }

    for _, test := range tests {
        t.Run(test.url+test.inputStr, func(t *testing.T) {
            response, err := makeRequest(test.url, test.inputStr)
            if err != nil {
                t.Fatalf("Expected no error, got %v", err)
            }
            if response != test.expected {
                t.Fatalf("Expected %v, got %v", test.expected, response)
            }
        })
    }
}