package main

import (
	"Golang_HTTP_Server/api"

	_ "modernc.org/sqlite"
)

func main() {
	server := api.NewServer()
	server.Run(":8080", "database/entries.db", "sqlite", "database/schema.sql")
}

// Database handling: Your initDB function in api/server.go is responsible for initializing the database connection and executing the SQL script from the initfile.
// 					This is a good approach, but you might want to consider separating the database initialization logic into its own package or file to keep the server logic clean.

// Error handling: In your initDB function, you're using log.Fatalf to handle errors.
// 					While this is acceptable for development, you might want to consider using a more robust error handling mechanism,
// 					such as returning errors and letting the caller decide how to handle them.

// Imports: You're using some external libraries (e.g., github.com/ayden-boyko/Convert_Service_Go/pkg), which is fine.
// 		However, you might want to consider vendoring these dependencies to ensure your project is self-contained.
