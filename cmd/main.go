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

func main() {
	server := api.NewHTTPServer()
	// runs and checks for errors

	go func() {
		// TODO: PASS CONTEXT TO HANDLERS
		if err := server.Run(":8080", "database/entries.db", "sqlite", "database/schema.sql"); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error running server: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.SafeShutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	log.Println("HTTP server shutdown safely completed")
}

// TODO: add an api rate limiter (middleware)?

// Database handling: Your initDB function in api/server.go is responsible for initializing the database connection and executing the SQL script from the initfile.
// 					This is a good approach, but you might want to consider separating the database initialization logic into its own package or file to keep the server logic clean.

// Error handling: In your initDB function, you're using log.Fatalf to handle errors.
// 					While this is acceptable for development, you might want to consider using a more robust error handling mechanism,
// 					such as returning errors and letting the caller decide how to handle them.

// Imports: You're using some external libraries (e.g., github.com/ayden-boyko/Convert_Service_Go/pkg), which is fine.
// 		However, you might want to consider vendoring these dependencies to ensure your project is self-contained.
