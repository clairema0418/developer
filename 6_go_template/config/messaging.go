package config

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"os"
)

func NewWatermillLogger() watermill.LoggerAdapter {
	return watermill.NewStdLogger(false, false)
}

func NewPublisher() (message.Publisher, error) {
	logger := NewWatermillLogger()
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	return pubSub, nil
}

func NewSubscriber() (message.Subscriber, error) {
	logger := NewWatermillLogger()
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	return pubSub, nil
}

func GetTopic(topic string) string {
	prefix := os.Getenv("WATERMILL_TOPIC_PREFIX")
	return prefix + topic
}