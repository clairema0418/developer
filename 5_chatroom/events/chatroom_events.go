package events

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"chatroom/models"
)

const (
	ChatRoomCreated = "chatroom_created"
)

type ChatRoomCreatedPayload struct {
	ChatRoom models.ChatRoom
}

func NewChatRoomCreatedEvent(chatRoom models.ChatRoom) *message.Message {
	event := message.NewMessage(uuid.New().String(), ChatRoomCreatedPayload{ChatRoom: chatRoom})
	event.Metadata.Set("type", ChatRoomCreated)
	return event
}

func IsChatRoomCreatedEvent(msg *message.Message) bool {
	return msg.Metadata.Get("type") == ChatRoomCreated
}

func UnmarshalChatRoomCreatedEvent(msg *message.Message) (*ChatRoomCreatedPayload, error) {
	payload := &ChatRoomCreatedPayload{}
	err := msg.Unmarshal(payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}