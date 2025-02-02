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
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Server is a thin wrapper around http.ServeMux.
type Server struct {
	Router *http.ServeMux
	db     *sql.DB
	Cache  map[string]string // TODO set up cache
	// Cache implementation: You have a TODO comment about setting up caches.
	// 						You might want to consider using a library like github.com/patrickmn/go-cache or github.com/bradfitz/gomemcache
}

// NewServer creates a new server with an empty request multiplexer.
func NewServer() *Server {
	return &Server{
		Router: http.NewServeMux(),
		db:     nil,
	}
}

func (s *Server) initDB(db string, dbdriver string, initfile string) {
	var err error
	//fmt.Println("Initializing database...", db, dbdriver, initfile)
	s.db, err = sql.Open(dbdriver, db)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	fmt.Println("Database opened")

	// Read the contents of the initfile
	sqlScript, err := os.ReadFile(initfile)
	if err != nil {
		log.Fatalf("Error reading init file: %v", err)
	}

	_, err = s.db.Exec(string(sqlScript))
	if err != nil {
		log.Fatalf("Error initializing database: %v, error within %s", err, initfile)
	}
	fmt.Println("Database initialized")

	vals, err := s.db.Query("SELECT * FROM entries")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	fmt.Println("Database queried", vals)

}

// Run runs the server on the given port.
func (s *Server) Run(port string, db string, dbdriver string, initfile string) error {
	println("Server running on port " + port)
	s.initDB(db, dbdriver, initfile)
	return http.ListenAndServe(port, s.Router)
}
