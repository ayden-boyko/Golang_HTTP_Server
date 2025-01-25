package api

import (
	H "Golang_HTTP_Server/internal/handlers"
	"net/http"
)

func (s *Server) RegisterRoutes() {
	s.router.HandleFunc("/", H.Home) // or http instead of s.router
	// s.router.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
	// 	if err := H.HandleShorten; err != nil {
	// 		log.Printf("Error in Home handler: %v", err)
	// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	}
	// })
	s.router.HandleFunc("/shorten", H.HandleShorten)
	s.router.HandleFunc("/{short_url}", H.HandleURL)
	s.router.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("website/styles"))))
	s.router.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("website/scripts"))))
}
