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
	"time"

	"github.com/patrickmn/go-cache"
)

// Server is a thin wrapper around http.ServeMux.
type Server struct {
	Router *http.ServeMux
	db     *sql.DB
	Cache  *cache.Cache
}

// NewServer creates a new server with an empty request multiplexer.
func NewServer() *Server {
	return &Server{
		Router: http.NewServeMux(),
		db:     nil,
		Cache:  cache.New(10*time.Minute, 10*time.Minute), // defualt expo, purge expo (10,10 minutes)
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

	_, err = s.db.Query("SELECT * FROM entries")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	//fmt.Println("Database queried", vals)

}

// Run runs the server on the given port.
func (s *Server) Run(port string, db string, dbdriver string, initfile string) error {
	println("Server running on port " + port)
	s.initDB(db, dbdriver, initfile)
	s.RegisterRoutes()
	return http.ListenAndServe(port, s.Router)
}

func (s *Server) checkCache(next http.Handler) http.Handler { // TODO save entry into cache, should be in a goroutine and a separate function? MIDDLEWARE
	/// on get resquest, check if entry is in cache, on post request, add entry to cache
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry, found := s.Cache.Get(r.URL.Path[1:])
		if found {
			fmt.Println("Cache hit")
			w.Write(entry.([]byte))
			fmt.Println(entry.([]byte))
			return
		}
		fmt.Println("Cache miss")
		next.ServeHTTP(w, r)
	})
}

func (s *Server) LogAction(next http.Handler) http.Handler { // TODO set up logs USE .log files
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Request completed: %s %s", r.Method, r.URL.Path)
	})
}
