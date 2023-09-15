package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
)

var (
	Publisher message.Publisher
)

func init() {
	logger := watermill.NewStdLogger(false, false)

	pub, err := gochannel.NewGoChannel(gochannel.Config{}, logger)
	if err != nil {
		log.Fatalf("Failed to create publisher: %v", err)
	}

	Publisher = pub
}

func PublishYouTubeVideoCreated(videoID string) error {
	msg := message.NewMessage(watermill.NewUUID(), []byte(videoID))
	return Publisher.Publish("youtube_video_created", msg)
}

func PublishYouTubeAudioExtracted(videoID string) error {
	msg := message.NewMessage(watermill.NewUUID(), []byte(videoID))
	return Publisher.Publish("youtube_audio_extracted", msg)
}

func PublishYouTubeTranscriptionCompleted(videoID string) error {
	msg := message.NewMessage(watermill.NewUUID(), []byte(videoID))
	return Publisher.Publish("youtube_transcription_completed", msg)
}