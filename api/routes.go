package api

import (
	H "Golang_HTTP_Server/internal/handlers"
	internal "Golang_HTTP_Server/internal/models"
	"log"
	"net/http"

	bvr "github.com/ayden-boyko/Log_Service_Go/pkg"
	"github.com/patrickmn/go-cache"
)

func (s *HTTPServer) RegisterRoutes() {
	// Create Data manager
	manager, err := internal.NewDataManager(s.db)
	if err != nil {
		log.Printf("Error creating manager: %v", err)
	}
	log.Printf("Manager created, %v", manager)

	// Create Logger
	Log_MWare, err := bvr.NewBeaverFromFile("./configs/logger.json")
	if err != nil {
		log.Printf("Error creating logger: %v", err)
	}

	// Home route
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := H.Home(w, r); err != nil {
			log.Printf("Error in Home handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// /shorten route with logging middleware
	shortenHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := H.HandleShorten(w, r, manager); err != nil {
			log.Printf("Error in HandleShorten handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
	s.Router.Handle("/shorten", bvr.LoggingMiddleware(Log_MWare, shortenHandler))

	// /{short_url} route with logging and caching middleware
	urlHandler := s.checkCache(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var val string
		if val, err = H.HandleURL(w, r, manager); err != nil {
			log.Printf("Error in HandleURL handler: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			s.cache.Add(r.URL.Path[1:], val, cache.DefaultExpiration)
		}
	}))
	s.Router.Handle("/{short_url}", bvr.LoggingMiddleware(Log_MWare, urlHandler))

	// Static file routes
	s.Router.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("website/styles"))))
	s.Router.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("website/scripts"))))
}
