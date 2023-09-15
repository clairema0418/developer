package message_publisher

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
)

var (
	publisher message.Publisher
)

func init() {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
	publisher = pubSub
}

func PublishQuestionGenerationCompleted(questionCount int) {
	msg := message.NewMessage(watermill.NewUUID(), []byte(questionCount))
	err := publisher.Publish("questionGenerationCompleted", msg)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
}