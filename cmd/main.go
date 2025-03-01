package main

import (
	"Golang_HTTP_Server/api"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "modernc.org/sqlite"
)

// main is the entry point for the program.
func main() {
	// Create a new HTTP server. This server will be responsible for running the
	// API and handling requests.
	server := api.NewHTTPServer()

	// Run the server in a separate goroutine. This allows the server to run
	// concurrently with the other code.
	go func() {
		// Run the server and check for errors. This will block until the server
		// is shutdown.
		if err := server.Run(":8080", "database/entries.db", "sqlite", "database/schema.sql"); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error running server: %v", err)
		}
	}()

	// Create a channel to receive signals. This will allow us to gracefully
	// shutdown the server when it receives a SIGINT or SIGTERM.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal to be received.
	<-sigChan

	// Create a context with a timeout to shut down the server.
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	// Shutdown the server. This will block until the server is shutdown.
	if err := server.SafeShutdown(shutdownCtx); err != nil {
		log.Fatalf("\n HTTP server shutdown failed: %v", err)
	}
	log.Println("\n HTTP server shutdown safely completed")
}

// TODO NGINX?

// Imports: You're using some external libraries (e.g., github.com/ayden-boyko/Convert_Service_Go/pkg), which is fine.
// 		However, you might want to consider vendoring these dependencies to ensure your project is self-contained.
