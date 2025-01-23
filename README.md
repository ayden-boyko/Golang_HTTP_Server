# Golang_HTTP_Server

A HTTP server dedicated to storing tiny and redirecting URLs

## Overview

This repository contains a simple HTTP server that is capable of storing and redirecting URLs. It uses SQLite as a database to store the URLs and a custom URL shortener module I created to shorten the URLs.

## Features

- Stores URLs in a SQLite database
- Redirects URLs to their original URL
- Uses a custom URL shortener module to shorten URLs

## This is an overview of how this golang URL shortener works:

1. A URL is submitted
   1. This URL is given a UUID
   2. This UUID is converted to base62
   3. All of this (UUID, base62_id, URL) are stored in a SQLite database
2. The base62 id is appended to the end of the shortener URL (e.g., www.gourl.com/2XT7A)
3. When the URL is entered:
   1. The base62 id is unhashed
   2. The cache is checked
      1. If the entry is in the cache, redirect to the respective link
      2. If not in the cache, retrieve the link from the database and add it to the cache
4. The link is then served with a 301 status code

## How to use

- Clone the repository
- Run the server using `go run main.go`
- The server will be accessible on `http://localhost:8080`
- To shorten a URL, use the following format: `http://localhost:8080/shorten?url=<url to shorten>`
- To redirect a URL, use the following format: `http://localhost:8080/<shortened url>`
