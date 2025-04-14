package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kiznlh/lettered-backend/internal/storage"
	"github.com/kiznlh/lettered-backend/internal/ws"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	storage.Connect()
	hub := ws.NewHub()
	go hub.Run()

	router := http.NewServeMux()
	router.HandleFunc("/", serveHome)
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)

	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port localhost:8080")
	server.ListenAndServe()
}
