package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for JSON response
type Response struct {
	Message string `json:"message"`
}

// Root handler
func handleRoot(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Welcome to Go Server!"}
	writeJSON(w, response)
}

// /hello route handler
func handleHello(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello from /hello route!"}
	writeJSON(w, response)
}

// /api/data route handler
func handleData(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"user":  "gopher",
		"email": "gopher@golang.org",
	}
	writeJSON(w, data)
}

// JSON response helper
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	// Routes
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/hello", handleHello)
	http.HandleFunc("/api/data", handleData)

	// Start server on port 8000
	port := ":8000"
	fmt.Printf("🚀 Server running at http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
