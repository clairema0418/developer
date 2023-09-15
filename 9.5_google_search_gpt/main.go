package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	http.HandleFunc("/ask", handleUserQuestion)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}