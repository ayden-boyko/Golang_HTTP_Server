package api

import (
	H "Golang_HTTP_Server/internal/handlers"
	internal "Golang_HTTP_Server/internal/models"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() {
	// Data manager created
	manager, err := internal.NewDataManager(s.db)
	if err != nil {
		log.Printf("Error creating manager: %v", err)
	}

	log.Printf("manager created, %v", manager)

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := H.Home(w, r); err != nil {
			log.Printf("Error in Home handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if err := H.Home; err != nil {
	// 		log.Printf("Error in Home handler: %v", err)
	// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	}
	// })

	// http.Handle("/", http.FileServer(http.Dir("./static")))
	s.Router.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if err := H.HandleShorten(w, r, manager); err != nil {
			log.Printf("Error in HandleShorten handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	s.Router.HandleFunc("/{short_url}", func(w http.ResponseWriter, r *http.Request) {
		if err := H.HandleURL(w, r, manager); err != nil {
			log.Printf("Error in HandleURL handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Then modify your routes:
	s.Router.Handle("/styles/", loggingMiddleware(http.StripPrefix("/styles/", http.FileServer(http.Dir("website/styles")))))
	s.Router.Handle("/scripts/", loggingMiddleware(http.StripPrefix("/scripts/", http.FileServer(http.Dir("website/scripts")))))

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Request completed: %s %s", r.Method, r.URL.Path)
	})
}
