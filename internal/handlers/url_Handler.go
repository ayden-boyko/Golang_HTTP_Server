package internal

//api for shortening urls used when html form is submited
import (
	"fmt"
	"log"
	"net/http"

	models "Golang_HTTP_Server/internal/models"

	pkg "github.com/ayden-boyko/Convert_Service_Go/pkg"
)

func HandleURL(w http.ResponseWriter, r *http.Request, dm *models.DataManagerImpl) error {

	switch r.Method {
	case "GET":
		Base62_id := r.URL.Path[1:]
		Base10_id, err := pkg.Base62ToUint64(Base62_id)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Base62_id:", Base62_id)
		fmt.Println("Base10_id:", Base10_id)

		val, err := dm.GetEntry(Base10_id)

		if err != nil {
			return err
		}

		w.Header().Set("Status", "200")

		// redirect to long url
		http.Redirect(w, r, val, http.StatusMovedPermanently)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	return nil
}
