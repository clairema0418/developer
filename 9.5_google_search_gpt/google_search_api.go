package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type GoogleSearchResult struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}

func googleSearch(query string) ([]GoogleSearchResult, error) {
	loadConfig()

	searchURL := GoogleSearchAPIEndpoint + "?key=" + GoogleSearchAPIKey + "&cx=" + os.Getenv("GOOGLE_CSE_ID") + "&q=" + url.QueryEscape(query)

	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Google Search API request failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse struct {
		Items []GoogleSearchResult `json:"items"`
	}

	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.Items, nil
}