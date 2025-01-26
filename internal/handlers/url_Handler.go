package internal

//api for shortening urls used when html form is submited
import (
	"fmt"
	"log"
	"net/http"

	pkg "github.com/ayden-boyko/Convert_Service_Go/pkg"
)

// TODO: Add error handling
// TODO: USE CODE 301 for redirecting, long url is cached, so if tiny url is entered, a request to this server isnt made
func HandleURL(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		Base62_id := r.URL.Path[1:]
		Base10_id, err := pkg.Base62ToUint64(Base62_id)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Base62_id:", Base62_id)
		fmt.Println("Base10_id:", Base10_id)

		// use base10_id to search cache and then sqlitedb if not found

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// save entry into sqlite db and/or cache, should be in a goroutine and a separate function?

}
