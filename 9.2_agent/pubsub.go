package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var pubsubConnection *gochannel.GoChannel

func setupPubSub() {
	pubsubConnection = gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
}

func publishMessage(topic string, payload []byte) error {
	msg := message.NewMessage(watermill.NewUUID(), payload)
	return pubsubConnection.Publish(topic, msg)
}

func subscribe(topic string, handler func(*message.Message) error) error {
	_, err := pubsubConnection.Subscribe(topic, handler)
	return err
}

func init() {
	setupPubSub()
}