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
