package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	GoogleSearchAPIKey = os.Getenv("GOOGLE_SEARCH_API_KEY")
	ChatGPTAPIKey = os.Getenv("CHATGPT_API_KEY")
	GoogleSearchAPIEndpoint = os.Getenv("GOOGLE_SEARCH_API_ENDPOINT")
	ChatGPTAPIEndpoint = os.Getenv("CHATGPT_API_ENDPOINT")
}

func httpGetRequest(url string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func parseGoogleSearchResults(data []byte) ([]GoogleSearchResult, error) {
	var results []GoogleSearchResult
	err := json.Unmarshal(data, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func parseChatGPTResponse(data []byte) (string, error) {
	var response ChatGPTResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		return "", err
	}
	return response.Choices[0].Text, nil
}