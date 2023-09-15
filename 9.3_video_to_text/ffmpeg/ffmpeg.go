package ffmpeg

import (
	"fmt"
	"os/exec"
)

func ExtractAudio(videoFile string, audioFile string) error {
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-vn", "-acodec", "pcm_s16le", "-ar", "16000", "-ac", "1", audioFile)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error extracting audio: %v", err)
	}
	return nil
}