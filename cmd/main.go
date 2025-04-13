package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port localhost:8080")
	server.ListenAndServe()
}
