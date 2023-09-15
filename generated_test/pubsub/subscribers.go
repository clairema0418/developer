package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"workers"
)

var (
	subscriber message.Subscriber
)

func init() {
	subscriber = createSubscriber()
}

func createSubscriber() message.Subscriber {
	subscriber, err := gochannel.NewGoChannel(
		gochannel.Config{
			Persistent: true,
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	return subscriber
}

func SubscribeToMessages() {
	subscribeToYouTubeVideoCreated()
	subscribeToYouTubeAudioExtracted()
	subscribeToYouTubeTranscriptionCompleted()
}

func subscribeToYouTubeVideoCreated() {
	messages, err := subscriber.Subscribe("youtube_video_created")
	if err != nil {
		panic(err)
	}

	go processYouTubeVideoCreated(messages)
}

func processYouTubeVideoCreated(messages <-chan *message.Message) {
	for msg := range messages {
		workers.ExtractYouTubeAudio(msg)
		msg.Ack()
	}
}

func subscribeToYouTubeAudioExtracted() {
	messages, err := subscriber.Subscribe("youtube_audio_extracted")
	if err != nil {
		panic(err)
	}

	go processYouTubeAudioExtracted(messages)
}

func processYouTubeAudioExtracted(messages <-chan *message.Message) {
	for msg := range messages {
		workers.TranscribeAudio(msg)
		msg.Ack()
	}
}

func subscribeToYouTubeTranscriptionCompleted() {
	messages, err := subscriber.Subscribe("youtube_transcription_completed")
	if err != nil {
		panic(err)
	}

	go processYouTubeTranscriptionCompleted(messages)
}

func processYouTubeTranscriptionCompleted(messages <-chan *message.Message) {
	for msg := range messages {
		workers.DisplayTranscription(msg)
		msg.Ack()
	}
}

func Middleware() *middleware.CorrelationID {
	return middleware.NewCorrelationID(subscriber.GenerateMessageID)
}