package main

import (
	"log"
	"net/http"
)

// more control over server
// look into server() params
func main() {
	s := &http.Server{
		Addr: ":8080",
	}

	log.Fatal(s.ListenAndServe())
}
