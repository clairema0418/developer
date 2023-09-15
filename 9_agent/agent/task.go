package agent

import (
	"sync"
	"audio_splitter"
	"audio_to_text_converter"
)

type AudioFile struct {
	FilePath string
	Chunks   []string
}

func processAudioFiles(audioFiles []AudioFile) ([]string, error) {
	var wg sync.WaitGroup
	texts := make([]string, 0)
	textsMutex := &sync.Mutex{}

	for _, audioFile := range audioFiles {
		wg.Add(1)
		go func(audioFile AudioFile) {
			defer wg.Done()
			chunks, err := audio_splitter.SplitAudioFile(audioFile.FilePath)
			if err != nil {
				return
			}

			for _, chunk := range chunks {
				text, err := audio_to_text_converter.ConvertAudioToText(chunk)
				if err != nil {
					continue
				}

				textsMutex.Lock()
				texts = append(texts, text)
				textsMutex.Unlock()
			}
		}(audioFile)
	}

	wg.Wait()
	return texts, nil
}