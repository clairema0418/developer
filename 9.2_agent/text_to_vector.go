package main

import (
	"errors"
	"strings"
)

type TextSegment struct {
	ID        uint
	AudioID   uint
	StartTime float64
	EndTime   float64
	Text      string
}

type VectorSegment struct {
	ID        uint
	TextID    uint
	Vector    []float64
}

func convertTextToVector(textSegment TextSegment) (VectorSegment, error) {
	if textSegment.Text == "" {
		return VectorSegment{}, errors.New("empty text segment")
	}

	words := strings.Split(textSegment.Text, " ")
	vector := make([]float64, len(words))

	for i, word := range words {
		vector[i] = wordToVector(word)
	}

	return VectorSegment{
		TextID: textSegment.ID,
		Vector: vector,
	}, nil
}

func wordToVector(word string) float64 {
	var vectorValue float64
	for _, char := range word {
		vectorValue += float64(char)
	}
	return vectorValue / float64(len(word))
}