package internal

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	models "Golang_HTTP_Server/internal/models"

	pkg "github.com/ayden-boyko/Convert_Service_Go/pkg"
	"github.com/google/uuid"
)

func HandleShorten(w http.ResponseWriter, r *http.Request, dm *models.DataManagerImpl) error {
	fmt.Println("received post request")
	switch r.Method {
	case "POST":

		var req models.ShortenRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return errors.New("error decoding request body")
		}

		uuid := uuid.New()

		base10_id := binary.BigEndian.Uint64(uuid[0:8]) & 0xFFFFFFFF // Use only lower 32 bits

		base62_id, err := pkg.Uint64ToBase62(base10_id)

		if err != nil {
			log.Fatal("error converting to base62", err)
		}

		entry := models.Entry{
			Id:           base10_id,
			Base62_id:    base62_id,
			LongUrl:      req.URL,
			Date_Created: time.Date(2025, 1, 26, 16, 11, 35, 0, time.FixedZone("EST", -5*60*60)),
		}

		fmt.Println("entry:", entry)

		var response interface{}
		var existing_base62_id string

		// TODO, make a goroutine to save entry into sqlitedb?
		// check if entry already exists
		if existing_base62_id, err := dm.PushData(entry); err != nil && err.Error() == "entry already exists" {
			fmt.Println("val:", existing_base62_id)

		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

		if existing_base62_id == "" {
			base62_id = existing_base62_id
		}

		// create the response

		tiny_url := "www.gourl.com/" + base62_id

		fmt.Println("tiny_url:", tiny_url)

		response = struct {
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

	return nil
}
