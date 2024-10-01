package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Connect to the SSE server
	resp, err := http.Get("http://localhost:8888/events")
	if err != nil {
		fmt.Printf("Error connecting to server: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Make sure the server sent a proper Content-Type header
	if ct := resp.Header.Get("Content-Type"); ct != "text/event-stream" {
		fmt.Printf("Expected Content-Type text/event-stream but got %s\n", ct)
		return
	}

	// Read events from the response body
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 5 && line[:5] == "data:" {
			data := line[6:]
			fmt.Printf("Received event: %s\n", data)
		}
	}

	// Check for errors in the scanner
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading from server: %v\n", err)
	}

	fmt.Println("Connection closed")
}
