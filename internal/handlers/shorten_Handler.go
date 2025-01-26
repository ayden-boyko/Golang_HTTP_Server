package internal

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	models "Golang_HTTP_Server/internal/models"

	pkg "github.com/ayden-boyko/Convert_Service_Go/pkg"
	"github.com/google/uuid"
)

func HandleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received post request")
	switch r.Method {
	case "POST":

		var req models.ShortenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		uuid := uuid.New()

		base10_id := binary.BigEndian.Uint64(uuid[:8])

		base62_id, err := pkg.Uint64ToBase62(base10_id)

		if err != nil {
			log.Fatal(err)
		}

		tiny_url := "www.gourl.com/" + base62_id

		entry := models.Entry{
			Id:           base10_id,
			Base62_id:    base62_id,
			LongUrl:      req.URL,
			Date_Created: time.Date(2025, 1, 26, 16, 11, 35, 0, time.FixedZone("EST", -5*60*60)),
		}

		// save entry into sqlite db and/or cache, should be in a goroutine and a separate function?
		fmt.Println("entry:", entry)

		response := struct {
			ShortUrl string `json:"short_url"`
		}{
			ShortUrl: tiny_url,
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("response:", response)
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
