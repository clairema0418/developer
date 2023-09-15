package events

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"chatroom/models"
)

const (
	MessageSent = "message_sent"
)

type MessageSentPayload struct {
	Message models.Message
}

func NewMessageSentEvent(message models.Message) *message.Message {
	payload := MessageSentPayload{Message: message}
	event := message.NewMessage(MessageSent, payload)
	return event
}

func GetMessageSentPayload(msg *message.Message) (*MessageSentPayload, error) {
	payload := &MessageSentPayload{}
	err := msg.UnmarshalPayload(payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func SetupMessageEvents(pubSub *gochannel.GoChannel) {
	pubSub.AddNoPublisherHandler(MessageSent, func(msg *message.Message) error {
		payload, err := GetMessageSentPayload(msg)
		if err != nil {
			return err
		}
		handleMessageSent(payload)
		return nil
	})
}

func handleMessageSent(payload *MessageSentPayload) {
	// Handle the message sent event here
}