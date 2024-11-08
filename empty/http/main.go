package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define the handler for GET /v1/get_empty_connection
	http.HandleFunc("/v1/get_empty_connection", func(w http.ResponseWriter, r *http.Request) {
		// Set the response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Create the response data
		response := map[string]string{
			"message": "Connection is empty",
		}

		// Encode the response as JSON and send it
		json.NewEncoder(w).Encode(response)
	})

	// Run the server on port 8888
	http.ListenAndServe(":8888", nil)
}
