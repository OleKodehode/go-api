package main

import (
	"log"
	"net/http"

	"github.com/OleKodehode/december-api-go/internal/handler"
)

func main () {
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("GET /health", handler.Health)

	log.Println("Server starting on https://localhost:8001")
	log.Fatal(http.ListenAndServeTLS(":8001", "localhost.crt", "localhost.key", mux))
}