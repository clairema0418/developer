package search

import (
	"fmt"
	"strings"

	"github.com/yourusername/yourproject/audio_extractor"
	"github.com/yourusername/yourproject/azure_tts"
	"github.com/yourusername/yourproject/ffmpeg"
	"github.com/yourusername/yourproject/text_segmenter"
	"github.com/yourusername/yourproject/time_recorder"
)

func SearchVideoForText(videoPath, searchText string) {
	audioPath := ffmpeg.ExtractAudio(videoPath)
	audioText := azure_tts.ConvertAudioToText(audioPath)
	segments := text_segmenter.SegmentText(audioText)
	timestamps := time_recorder.RecordTimestamps(segments)

	searchText = strings.ToLower(searchText)
	for i, segment := range segments {
		if strings.Contains(strings.ToLower(segment), searchText) {
			fmt.Printf("Found text '%s' in segment %d at time %s\n", searchText, i+1, timestamps[i])
		}
	}
}