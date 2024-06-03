package main

import (
	"fmt"
	"net/http"
)

// healthzHandler responds with a 200 OK status to indicate the service is healthy.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func main() {
	// Register the /healtz endpoint with the healthzHandler.
	http.HandleFunc("/healtz", healthzHandler)

	// Start the HTTP server on port 8080.
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
