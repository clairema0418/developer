package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	GoogleSearchAPIKey    string
	ChatGPTAPIKey         string
	GoogleSearchAPIEndpoint string
	ChatGPTAPIEndpoint    string
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GoogleSearchAPIKey = os.Getenv("GOOGLE_SEARCH_API_KEY")
	ChatGPTAPIKey = os.Getenv("CHATGPT_API_KEY")
	GoogleSearchAPIEndpoint = "https://www.googleapis.com/customsearch/v1"
	ChatGPTAPIEndpoint = "https://api.openai.com/v1/engines/davinci-codex/completions"
}