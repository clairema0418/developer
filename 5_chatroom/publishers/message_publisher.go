package publishers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"chatroom_backend/events"
	"chatroom_backend/models"
)

type MessagePublisher struct {
	pub *gochannel.PubSub
}

func NewMessagePublisher(pub *gochannel.PubSub) *MessagePublisher {
	return &MessagePublisher{
		pub: pub,
	}
}

func (mp *MessagePublisher) PublishMessageSent(messageSent *models.Message) error {
	messageSentEvent := events.MessageSent{
		ID:        messageSent.ID,
		ChatRoomID: messageSent.ChatRoomID,
		UserID:    messageSent.UserID,
		Content:   messageSent.Content,
		CreatedAt: messageSent.CreatedAt,
	}

	messageSentJSON, err := json.Marshal(messageSentEvent)
	if err != nil {
		return err
	}

	msg := message.NewMessage(messageSent.ID, messageSentJSON)
	return mp.pub.Publish(events.MessageSentTopic, msg)
}