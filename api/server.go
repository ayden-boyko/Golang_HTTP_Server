// Package api provides a simple HTTP server for handling requests.
//
// It includes a function for creating a new server, registering routes
// with the server, and running the server.
//
// The server is an http.Handler, so it can be passed to http.ListenAndServe
// or used in an http.Server.
//
// The server is a thin wrapper around http.ServeMux, so it supports all
// the same methods as http.ServeMux.
package api

import (
	"net/http"
)

// Server is a thin wrapper around http.ServeMux.
type Server struct {
	router *http.ServeMux
}

// NewServer creates a new server with an empty request multiplexer.
func NewServer() *Server {
	return &Server{
		router: http.NewServeMux(),
	}
}

// Run runs the server on the given port.
func (s *Server) Run(port string) error {
	println("Server running on port " + port)
	return http.ListenAndServe(port, s.router)
}
