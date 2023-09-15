package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func DownloadYouTubeAudio(videoURL string) (string, error) {
	videoID := extractVideoID(videoURL)
	if videoID == "" {
		return "", fmt.Errorf("invalid YouTube URL")
	}

	audioFileName := fmt.Sprintf("%s.mp3", videoID)
	audioFilePath := fmt.Sprintf("audio_files/%s", audioFileName)

	err := os.MkdirAll("audio_files", os.ModePerm)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", "-o", audioFilePath, videoURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("youtube-dl error: %s", output)
		return "", err
	}

	return audioFilePath, nil
}

func extractVideoID(videoURL string) string {
	parts := strings.Split(videoURL, "watch?v=")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}