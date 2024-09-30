package main

import (
	"log"
	"net/http"
)

func main() {
	// Define the directory to serve files from (current directory)
	dir := "./"

	// Create a file server handler that serves files from the specified directory
	fs := http.FileServer(http.Dir(dir))

	// Serve static files on the root URL "/"
	http.Handle("/", fs)

	// Define the port to listen on (e.g., 8080)
	port := ":8080"
	log.Printf("Serving files from %s on http://localhost%s\n", dir, port)

	// Start the HTTP server
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
