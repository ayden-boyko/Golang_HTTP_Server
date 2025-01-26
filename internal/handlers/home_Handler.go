package internal

import (
	"log"
	"net/http"
)

// the api route for the home page
// returns the website/main.html

func Home(w http.ResponseWriter, r *http.Request) error {

	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	switch r.URL.Path {

	case "/":
		http.ServeFile(w, r, "website/main.html")
	case "images/favicon.ico":
		http.ServeFile(w, r, "website/images/favicon.ico")
	case "styles/style.css":
		http.ServeFile(w, r, "website/styles/style.css")
	case "scripts/script.js":
		http.ServeFile(w, r, "website/scripts/script.js")

	default:
		http.NotFound(w, r)
	}

	return nil
}
