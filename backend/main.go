package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	data := readCsv()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, "Failed to encode data as JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	http.ListenAndServe(":8080", addCORSHeaders(http.DefaultServeMux))
}

func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow requests from http://localhost:5500
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If the request is an OPTIONS preflight request, return immediately
		if r.Method == "OPTIONS" {
			return
		}

		// Call the original handler
		handler.ServeHTTP(w, r)
	})
}
