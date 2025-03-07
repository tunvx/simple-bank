package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// Define the handler for GET /v1/test/empty_get
	http.HandleFunc("/v1/test/empty_get", func(w http.ResponseWriter, r *http.Request) {
		// Set the response header to JSON
		w.Header().Set("Content-Type", "application/json")

		// Create the response data
		response := map[string]string{
			"message": "Ok",
		}

		// Encode the response as JSON and send it
		json.NewEncoder(w).Encode(response)
	})

	// Run the server on port 8083
	http.ListenAndServe(":8083", nil)
}
