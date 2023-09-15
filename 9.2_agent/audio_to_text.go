package main

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type TextSegment struct {
	ID          uint
	AudioSegmentID uint
	Text        string
}

func convertAudioToText(audioSegment AudioSegment) (TextSegment, error) {
	audioFilePath := filepath.Join("audio_segments", audioSegment.FileName)
	textFilePath := strings.TrimSuffix(audioFilePath, filepath.Ext(audioFilePath)) + ".txt"

	cmd := exec.Command("ffmpeg", "-i", audioFilePath, "-vn", "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1", "temp.wav")
	err := cmd.Run()
	if err != nil {
		return TextSegment{}, err
	}

	cmd = exec.Command("deepspeech", "--model", "deepspeech-0.9.3-models.pbmm", "--scorer", "deepspeech-0.9.3-models.scorer", "--audio", "temp.wav")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return TextSegment{}, err
	}

	err = ioutil.WriteFile(textFilePath, output, 0644)
	if err != nil {
		return TextSegment{}, err
	}

	err = os.Remove("temp.wav")
	if err != nil {
		return TextSegment{}, err
	}

	text, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		return TextSegment{}, err
	}

	if len(text) == 0 {
		return TextSegment{}, errors.New("empty text")
	}

	textSegment := TextSegment{
		AudioSegmentID: audioSegment.ID,
		Text:        string(text),
	}

	return textSegment, nil
}