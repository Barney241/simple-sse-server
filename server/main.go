package main

import (
	"fmt"
	"net/http"
	"time"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins. You may want to restrict this to specific origins in a production environment.
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (you can replace this with real data)
	for i := 0; i < 10; i++ {
		event := fmt.Sprintf("Event %d", i)

		fmt.Fprintf(w, "data: %s\n\n", event)
		w.(http.Flusher).Flush()

		fmt.Println("Sent event:", event)

		time.Sleep(2 * time.Second)
	}

	// Simulate closing the connection
	// Comment this line to keep the connection open
	r.Context().Done()

	fmt.Println("Finished and closing the connection...")
}

func main() {
	http.HandleFunc("/events", eventsHandler)

	fmt.Println("Server is listening on port 8888...")
	http.ListenAndServe(":8888", nil)
}
