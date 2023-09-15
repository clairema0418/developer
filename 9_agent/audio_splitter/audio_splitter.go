package audio_splitter

import (
	"errors"
	"fmt"
	"os"
)

type AudioFile struct {
	FilePath string
	Chunks   []string
}

func splitAudioFile(filePath string) ([]string, error) {
	if filePath == "" {
		return nil, errors.New("file path cannot be empty")
	}

	// Placeholder implementation for splitting the audio file into chunks
	// Replace this with the actual implementation using an audio processing library
	chunks := []string{
		fmt.Sprintf("%s_chunk1", filePath),
		fmt.Sprintf("%s_chunk2", filePath),
		fmt.Sprintf("%s_chunk3", filePath),
	}

	return chunks, nil
}

func checkFileExists(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New("file does not exist")
	}
	return nil
}