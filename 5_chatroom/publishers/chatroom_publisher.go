package publishers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"chatroom_backend/events"
	"chatroom_backend/models"
)

type ChatRoomPublisher struct {
	pub *gochannel.GoChannel
}

func NewChatRoomPublisher(pub *gochannel.GoChannel) *ChatRoomPublisher {
	return &ChatRoomPublisher{
		pub: pub,
	}
}

func (cp *ChatRoomPublisher) PublishChatRoomCreated(chatRoom *models.ChatRoom) error {
	chatRoomCreatedEvent := events.ChatRoomCreated{
		ID:        chatRoom.ID,
		Name:      chatRoom.Name,
		IsPrivate: chatRoom.IsPrivate,
	}

	payload, err := json.Marshal(chatRoomCreatedEvent)
	if err != nil {
		return err
	}

	msg := message.NewMessage(chatRoom.ID, payload)
	return cp.pub.Publish(events.ChatRoomCreatedTopic, msg)
}