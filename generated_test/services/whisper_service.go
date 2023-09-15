package services

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

const whisperAPIURL = "https://api.whisper.com/v1/transcribe"

type WhisperService struct{}

func NewWhisperService() *WhisperService {
	return &WhisperService{}
}

func (ws *WhisperService) TranscribeAudio(audioFilePath string) (string, error) {
	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to open audio file")
	}
	defer audioFile.Close()

	req, err := http.NewRequest("POST", whisperAPIURL, audioFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to create request")
	}

	req.Header.Set("Content-Type", "audio/wav")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("WHISPER_API_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("Whisper API returned status code %d", resp.StatusCode)
	}

	transcriptionBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	return string(transcriptionBytes), nil
}