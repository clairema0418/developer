package workers

import (
	"encoding/json"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"github.com/yourusername/yourprojectname/pubsub"
	"github.com/yourusername/yourprojectname/services"
	"github.com/yourusername/yourprojectname/models"
)

type TranscriptionWorker struct {
	subscriber *gochannel.Subscriber
	whisperSvc *services.WhisperService
}

func NewTranscriptionWorker(subscriber *gochannel.Subscriber, whisperSvc *services.WhisperService) *TranscriptionWorker {
	return &TranscriptionWorker{
		subscriber: subscriber,
		whisperSvc: whisperSvc,
	}
}

func (tw *TranscriptionWorker) Start() {
	messages, err := tw.subscriber.Subscribe(pubsub.YoutubeAudioExtracted)
	if err != nil {
		log.Fatalf("Failed to subscribe to %s: %v", pubsub.YoutubeAudioExtracted, err)
	}

	go func() {
		for msg := range messages {
			var video models.YouTubeVideo
			err := json.Unmarshal(msg.Payload, &video)
			if err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				msg.Ack()
				continue
			}

			err = tw.processTranscription(&video)
			if err != nil {
				log.Printf("Failed to process transcription: %v", err)
			}

			msg.Ack()
		}
	}()
}

func (tw *TranscriptionWorker) processTranscription(video *models.YouTubeVideo) error {
	transcription, err := tw.whisperSvc.TranscribeAudio(video.AudioFilePath)
	if err != nil {
		return err
	}

	video.Transcription = transcription
	pubsub.PublishYoutubeTranscriptionCompleted(video)

	return nil
}