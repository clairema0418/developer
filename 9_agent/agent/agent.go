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

func createAgent() *Agent {
	return &Agent{
		wg: &sync.WaitGroup{},
	}
}

type Agent struct {
	wg *sync.WaitGroup
}

func (a *Agent) processAudioFiles(audioFiles []AudioFile) ([]string, error) {
	results := make([]string, 0)
	a.wg.Add(len(audioFiles))

	for _, audioFile := range audioFiles {
		go func(file AudioFile) {
			defer a.wg.Done()
			chunks, err := audio_splitter.SplitAudioFile(file.FilePath)
			if err != nil {
				return
			}

			file.Chunks = chunks
			for _, chunk := range file.Chunks {
				text, err := audio_to_text_converter.ConvertAudioToText(chunk)
				if err != nil {
					return
				}
				results = append(results, text)
			}
		}(audioFile)
	}

	a.wg.Wait()
	return results, nil
}