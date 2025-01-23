package api

//api for shortening urls
import (
	"fmt"
	"log"
	"net/http"

	pkg "github.com/ayden-boyko/Golang-URL-shrtnr/pkg"
)

// TODO: Add error handling
// TODO: MAKE BETTER
func Shorten(w http.ResponseWriter, r *http.Request) (uint64, string, string, error) {
	id, base62_id, short_url, err := pkg.Url_shortener(w.Header().Get("Origin"), r.FormValue("url"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortened URL:", short_url)
	return id, base62_id, short_url, nil
}
