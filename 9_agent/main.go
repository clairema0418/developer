package main

import (
	"fmt"
	"os"
	"agent"
	"audio_splitter"
	"audio_to_text_converter"
)

type AudioFile struct {
	FilePath string
	Chunks   []string
}

func main() {
	audioFiles := []AudioFile{
		{FilePath: "path/to/audio/file1.mp3"},
		{FilePath: "path/to/audio/file2.mp3"},
	}

	results, err := processAudioFiles(audioFiles)
	if err != nil {
		fmt.Println("Error processing audio files:", err)
		os.Exit(1)
	}

	fmt.Println("Text results:", results)
}

func processAudioFiles(audioFiles []AudioFile) ([]string, error) {
	agent := agent.CreateAgent()

	for _, audioFile := range audioFiles {
		chunks, err := audio_splitter.SplitAudioFile(audioFile.FilePath)
		if err != nil {
			return nil, err
		}

		for _, chunk := range chunks {
			agent.AddTask(func() {
				text, err := audio_to_text_converter.ConvertAudioToText(chunk)
				if err != nil {
					fmt.Println("Error converting audio to text:", err)
					return
				}

				audioFile.Chunks = append(audioFile.Chunks, text)
			})
		}
	}

	agent.Wait()

	var results []string
	for _, audioFile := range audioFiles {
		results = append(results, audioFile.Chunks...)
	}

	return results, nil
}