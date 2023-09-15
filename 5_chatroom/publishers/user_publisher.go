package publishers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"chatroom_backend/events"
)

type UserPublisher struct {
	Publisher *gochannel.Publisher
}

func NewUserPublisher(pub *gochannel.Publisher) *UserPublisher {
	return &UserPublisher{
		Publisher: pub,
	}
}

func (up *UserPublisher) PublishUserJoinedEvent(userJoinedEvent events.UserJoinedEvent) error {
	payload, err := json.Marshal(userJoinedEvent)
	if err != nil {
		return err
	}

	msg := message.NewMessage(userJoinedEvent.ID, payload)
	return up.Publisher.Publish(events.UserJoined, msg)
}

func (up *UserPublisher) PublishUserLeftEvent(userLeftEvent events.UserLeftEvent) error {
	payload, err := json.Marshal(userLeftEvent)
	if err != nil {
		return err
	}

	msg := message.NewMessage(userLeftEvent.ID, payload)
	return up.Publisher.Publish(events.UserLeft, msg)
}