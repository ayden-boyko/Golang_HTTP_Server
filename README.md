# Golang_HTTP_Server

A HTTP server dedicated to storing tiny and redirecting URLs

# Overview

This repository contains a simple HTTP server that is capable of storing and redirecting URLs. It uses SQLite as a database to store the URLs and a custom URL shortener module I created to shorten the URLs.

# Features

- Stores URLs in a SQLite database
- Redirects URLs to their original URL
- Uses a custom URL shortener module to shorten URLs

# How to use

- Clone the repository
- Run the server using `go run main.go`
- The server will be accessible on `http://localhost:8080`
- To shorten a URL, use the following format: `http://localhost:8080/shorten?url=<url to shorten>`
- To redirect a URL, use the following format: `http://localhost:8080/<shortened url>`
