package main

import (
	"Golang_HTTP_Server/api"
	"log"
	"net/http"
)

func main() {
	// Define routes
	http.HandleFunc("/", api.HandleHome)       // Serve the website
	http.HandleFunc("/{short_url}", api.Fetch) // Fetch and redirect to long URLs

	// Start server
	serverAddress := ":8080" // Default port
	log.Printf("Server is listening on %s", serverAddress)
	log.Fatal(http.ListenAndServe(serverAddress, nil)) // Start the server
}
