package main

import (
	"Golang_HTTP_Server/api"

	_ "modernc.org/sqlite"
)

func main() {
	server := api.NewServer()
	server.RegisterRoutes()
	server.Run(":8080", "database/entries.db", "sqlite", "database/schema.sql")
}

// TODO : implement a way for the handler to get and push data to the db, in an idiomatic way

// Database handling: Your initDB function in api/server.go is responsible for initializing the database connection and executing the SQL script from the initfile.
// 					This is a good approach, but you might want to consider separating the database initialization logic into its own package or file to keep the server logic clean.

// Error handling: In your initDB function, you're using log.Fatalf to handle errors.
// 					While this is acceptable for development, you might want to consider using a more robust error handling mechanism,
// 					such as returning errors and letting the caller decide how to handle them.

// Database connection: You're storing the database connection in the Server struct, which is a good approach.
// 						However, you might want to consider using a separate package for database operations to keep the server logic clean.

// Testing: You have some test files (tests/db_test.go and tests/cache_test.go) with TODO comments.
// 			You should prioritize writing tests for your database and cache logic to ensure they're working correctly.

// Package organization: Your package organization appears to be good,
// 					but you might want to consider reorganizing some of your packages to better reflect their responsibilities.
// 					For example, you could have a separate package for database operations, caching, and server logic.

// Imports: You're using some external libraries (e.g., github.com/ayden-boyko/Convert_Service_Go/pkg), which is fine.
// 		However, you might want to consider vendoring these dependencies to ensure your project is self-contained.
