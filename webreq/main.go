package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Basic GET request
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Reading response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("GET Response:", string(body))

	// POST request with JSON data
	data := map[string]string{
		"name":  "John Doe",
		"email": "john@example.com",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err = http.Post("https://jsonplaceholder.typicode.com/todos/1", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Custom HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Creating custom request
	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Adding headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer your-token-here")

	// Sending request using custom client
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Handling different status codes
	switch resp.StatusCode {
	case http.StatusOK:
		fmt.Println("Request successful")
	case http.StatusUnauthorized:
		fmt.Println("Unauthorized access")
	case http.StatusNotFound:
		fmt.Println("Resource not found")
	default:
		fmt.Printf("Unexpected status code: %d\n", resp.StatusCode)
	}

	// Example of concurrent requests using goroutines
	urls := []string{
		"https://api.example.com/1",
		"https://api.example.com/2",
		"https://api.example.com/3",
	}

	ch := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
				return
			}
			defer resp.Body.Close()
			ch <- fmt.Sprintf("Successfully fetched %s", url)
		}(url)
	}

	// Collecting results from goroutines
	for range urls {
		fmt.Println(<-ch)
	}
}

// Helper function to make HTTP requests with retry mechanism
func makeRequestWithRetry(url string, maxRetries int) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for i := 0; i < maxRetries; i++ {
		resp, err := client.Get(url)
		if err == nil {
			return resp, nil
		}

		// Wait before retrying (exponential backoff)
		time.Sleep(time.Duration(i+1) * time.Second)
	}

	return nil, fmt.Errorf("failed after %d retries", maxRetries)
}
