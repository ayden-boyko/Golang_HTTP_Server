package api

import (
	H "Golang_HTTP_Server/internal/handlers"
	"net/http"
)

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/", H.Home)
	//http.HandleFunc("/shorten", handlers.Shorten)
	http.HandleFunc("/{short_url}", H.Fetch)
}
