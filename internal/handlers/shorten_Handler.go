package internal

//api for shortening urls used when html form is submited
import (
	"fmt"
	"log"
	"net/http"

	Entry "Golang_HTTP_Server/internal/models"

	pkg "github.com/ayden-boyko/Golang-URL-shrtnr/pkg"
)

// TODO: Add error handling
// TODO: USE CODE 301 for redirecting, long url is cached, so if tiny url is entered, a request to this server isnt made
func Shorten(w http.ResponseWriter, r *http.Request) (Entry.Entry, error) {
	id, base62_id, long_url, err := pkg.Url_shortener(w.Header().Get("Origin"), r.FormValue("url"))

	if err != nil {
		log.Fatal(err)
	}

	entry := Entry.Entry{
		Id:        id,
		Base62_id: base62_id,
		LongUrl:   long_url,
	}

	// save entry into sqlite db and/or cache, should be in a goroutine and a separate function?

	fmt.Println("Entry Struct:", entry)
	return entry, nil
}
