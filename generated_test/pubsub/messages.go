package pubsub

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	YouTubeVideoCreatedTopic        = "youtube_video_created"
	YouTubeAudioExtractedTopic      = "youtube_audio_extracted"
	YouTubeTranscriptionCompletedTopic = "youtube_transcription_completed"
)

type YouTubeVideoCreatedMessage struct {
	ID  string
	URL string
}

type YouTubeAudioExtractedMessage struct {
	ID           string
	AudioFilePath string
}

type YouTubeTranscriptionCompletedMessage struct {
	ID            string
	Transcription string
}

func NewYouTubeVideoCreatedMessage(id, url string) *message.Message {
	msg := YouTubeVideoCreatedMessage{
		ID:  id,
		URL: url,
	}
	return message.NewMessage(id, msg)
}

func NewYouTubeAudioExtractedMessage(id, audioFilePath string) *message.Message {
	msg := YouTubeAudioExtractedMessage{
		ID:           id,
		AudioFilePath: audioFilePath,
	}
	return message.NewMessage(id, msg)
}

func NewYouTubeTranscriptionCompletedMessage(id, transcription string) *message.Message {
	msg := YouTubeTranscriptionCompletedMessage{
		ID:            id,
		Transcription: transcription,
	}
	return message.NewMessage(id, msg)
}