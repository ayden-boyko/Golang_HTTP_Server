// Package api provides a simple HTTP HTTPServer for handling requests.
//
// It includes a function for creating a new HTTPServer, registering routes
// with the HTTPServer, and running the HTTPServer.
//
// The HTTPServer is an http.Handler, so it can be passed to http.ListenAndServe
// or used in an http.HTTPServer.
//
// The HTTPServer is a thin wrapper around http.ServeMux, so it supports all
// the same methods as http.ServeMux.
package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

// HTTPServer is a thin wrapper around http.ServeMux.
type HTTPServer struct {
	Server *http.Server
	Router *http.ServeMux
	db     *sql.DB
	Cache  *cache.Cache
}

// NewHTTPServer creates a new HTTPServer with an empty request multiplexer.
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{},
		Router: http.NewServeMux(),
		db:     nil,
		Cache:  cache.New(10*time.Minute, 10*time.Minute), // default expo, purge expo (10,10 minutes)
	}
}

func (s *HTTPServer) initDB(db string, dbdriver string, initfile string) {
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

	rows, err := s.db.Query("SELECT * FROM entries")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	rows.Close()

	if err != nil {
		log.Fatalf("Error inserting into database: %v", err)
	}
}

// Run runs the HTTPServer on the given port.
func (s *HTTPServer) Run(port string, db string, dbdriver string, initfile string) error {
	println("HTTPServer running on port " + port)
	s.Server.Addr = port
	s.initDB(db, dbdriver, initfile)
	s.RegisterRoutes()
	s.Server.Handler = s.Router
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) checkCache(next http.HandlerFunc) http.HandlerFunc {
	// on get resquest, check if entry is in cache, on post request, add entry to cache

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		entry, found := s.Cache.Get(r.URL.Path[1:])

		fmt.Println("Checking cache: ", entry)

		if found {
			fmt.Println("Cache hit")
			w.Header().Set("Cache-Status", "HIT")
			w.Write(entry.([]byte))
			fmt.Println(entry.([]byte))
			return
		}
		fmt.Println("Cache miss")
		w.Header().Set("Cache-Status", "MISS")
		next.ServeHTTP(w, r)

		if w.Header().Get("Status") == "200" {
			go func() {
				s.Cache.Set(r.URL.Path[1:], w, cache.DefaultExpiration)
			}()
		}
	})
}

func (s *HTTPServer) SafeShutdown(ctx context.Context) error {
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}
