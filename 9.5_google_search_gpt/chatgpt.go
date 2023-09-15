package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ChatGPTRequest struct {
	Prompt string `json:"prompt"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func chatGPTQuery(prompt string) (string, error) {
	requestBody, err := json.Marshal(ChatGPTRequest{Prompt: prompt})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", ChatGPTAPIEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ChatGPTAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("ChatGPT API request failed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var chatGPTResponse ChatGPTResponse
	err = json.Unmarshal(body, &chatGPTResponse)
	if err != nil {
		return "", err
	}

	if len(chatGPTResponse.Choices) == 0 {
		return "", errors.New("No response from ChatGPT")
	}

	return chatGPTResponse.Choices[0].Text, nil
}