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
	cache  *cache.Cache
}

// NewHTTPServer creates a new HTTPServer with an empty request multiplexer.
//
// The HTTPServer uses a NewServeMux to handle requests, and a cache with a
// default expiration and purge time of 10 minutes.
func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{},
		Router: http.NewServeMux(),
		db:     nil,
		cache:  cache.New(10*time.Minute, 10*time.Minute), // 10 minutes
	}
}

// initDB initializes the database connection and executes the SQL script from the initfile.
//
// If the database connection is already open, it will be closed and reopened.
// If the initfile is empty, the database will not be initialized.
func (s *HTTPServer) initDB(db string, dbdriver string, initfile string) {
	var err error
	fmt.Println("Initializing database...", db, dbdriver, initfile)
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

	// Execute the SQL script to initialize the database
	_, err = s.db.Exec(string(sqlScript))
	if err != nil {
		log.Fatalf("Error initializing database: %v, error within %s", err, initfile)
	}
	fmt.Println("Database initialized")

	rows, err := s.db.Query("SELECT * FROM entries")
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	// Close the rows to release resources
	rows.Close()

	if err != nil {
		log.Fatalf("Error inserting into database: %v", err)
	}
}

// Run runs the HTTPServer on the given port.
//
// It first initializes the database connection and executes the SQL script
// from the initfile. If the database connection is already open, it will be
// closed and reopened. If the initfile is empty, the database will not be
// initialized.
//
// Then it registers the routes with the HTTPServer and runs it on the given
// port.
func (s *HTTPServer) Run(port string, db string, dbdriver string, initfile string) error {
	// Print the port number
	println("HTTPServer running on port " + port)
	s.Server.Addr = port // Set the port

	// Initialize the database
	s.initDB(db, dbdriver, initfile)

	// Register the routes
	s.RegisterRoutes()

	// Set the handler to the registered routes
	s.Server.Handler = s.Router

	// Run the server
	return s.Server.ListenAndServe()
}

// checkCache is a middleware function that checks the cache for an entry on GET requests
// and adds an entry to the cache on POST requests.
func (s *HTTPServer) checkCache(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the entry from the cache using the URL path as the key
		entry, found := s.cache.Get(r.URL.Path[1:])

		fmt.Println("Checking cache: ", r.URL.Path[1:])

		if found {
			// Cache hit: return the cached entry
			fmt.Println("Cache hit")
			w.Header().Set("Cache-Status", "HIT")
			w.Write(entry.([]byte)) // FIXME PANIC SERVING FIX THIS
			fmt.Println(entry.([]byte))
			return
		}

		// Cache miss: proceed to the next handler
		fmt.Println("Cache miss")
		w.Header().Set("Cache-Status", "MISS")
		next.ServeHTTP(w, r)

		// On successful response, add the entry to the cache
		if w.Header().Get("Status") == "200" {
			go func() {
				s.cache.Set(r.URL.Path[1:], w, cache.DefaultExpiration)
			}()
		}
	})
}

// SafeShutdown is a function that gracefully stops the server and closes the database connection.
func (s *HTTPServer) SafeShutdown(ctx context.Context) error {
	// Shutdown the server
	if err := s.Server.Shutdown(ctx); err != nil {
		return err
	}

	// Close the database connection
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}
