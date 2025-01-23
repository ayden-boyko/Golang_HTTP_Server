package api

import (
	"net/http"
)

// the api route for the home page
// returns the website/main.html

// TODO: Add error handling

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "website/main.html")
}
