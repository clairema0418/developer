package main

import (
	"fmt"
	"os"

	"./audio_extractor"
	"./azure_tts"
	"./ffmpeg"
	"./search"
	"./text_segmenter"
	"./time_recorder"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./main <video_file>")
		return
	}

	videoFile := os.Args[1]

	// Extract audio from video file
	audioFile, err := ffmpeg.ExtractAudio(videoFile)
	if err != nil {
		fmt.Println("Error extracting audio:", err)
		return
	}

	// Convert audio to text using Azure TTS
	text, err := azure_tts.AudioToText(audioFile)
	if err != nil {
		fmt.Println("Error converting audio to text:", err)
		return
	}

	// Segment text at punctuation marks
	segments := text_segmenter.SegmentText(text)

	// Record time for each segment
	timestamps, err := time_recorder.RecordTime(audioFile, segments)
	if err != nil {
		fmt.Println("Error recording time:", err)
		return
	}

	// Search for a specific segment of text
	query := "example text"
	result := search.Search(segments, timestamps, query)
	if result != nil {
		fmt.Printf("Found '%s' at %d seconds in the audio and video files.\n", query, result.Time)
	} else {
		fmt.Printf("'%s' not found in the audio and video files.\n", query)
	}
}