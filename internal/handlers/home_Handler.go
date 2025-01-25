package internal

import (
	"net/http"
)

// the api route for the home page
// returns the website/main.html

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, "website/main.html")
}
