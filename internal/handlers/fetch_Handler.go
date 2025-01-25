package internal

import (
	"net/http"
)

// This package contains the implementation for the /fetch route.
// It should take in a tiny url, decode the hash into its uint64 id, then use this id
// to search the sqlite db for the long url.

// TODO: implement functionality, check cache then db
func Fetch(w http.ResponseWriter, r *http.Request) {
	// Get the shortened URL from the request path
	short_url := r.URL.Path[1:] // Remove the leading slash
	println(short_url)
}
