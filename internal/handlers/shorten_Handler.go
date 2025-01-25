package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	Entry "Golang_HTTP_Server/internal/models"

	pkg "github.com/ayden-boyko/Golang-URL-shrtnr/pkg"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received post request")
	switch r.Method {
	case "POST":

		base10_id, base62_id, long_url, err := pkg.Url_shortener(w.Header().Get("Origin"), r.FormValue("url"))

		if err != nil {
			log.Fatal(err)
		}

		entry := Entry.Entry{
			Id:        base10_id,
			Base62_id: base62_id,
			LongUrl:   long_url,
		}

		// save entry into sqlite db and/or cache, should be in a goroutine and a separate function?
		fmt.Println("entry:", entry)

		response := struct {
			ShortUrl string `json:"short_url"`
		}{
			ShortUrl: entry.Base62_id,
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("response:", response)
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
