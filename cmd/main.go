package main

import (
	"Golang_HTTP_Server/api"
)

func main() {
	server := api.NewServer()
	server.RegisterRoutes()
	server.Run(":8080")
}
