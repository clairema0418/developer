package events

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
)

const (
	UserJoinedEvent = "user_joined"
	UserLeftEvent   = "user_left"
)

type UserJoinedPayload struct {
	UserID     uuid.UUID `json:"user_id"`
	ChatRoomID uuid.UUID `json:"chat_room_id"`
}

type UserLeftPayload struct {
	UserID     uuid.UUID `json:"user_id"`
	ChatRoomID uuid.UUID `json:"chat_room_id"`
}

func NewUserJoinedEvent(userID, chatRoomID uuid.UUID) *message.Message {
	payload := UserJoinedPayload{
		UserID:     userID,
		ChatRoomID: chatRoomID,
	}
	return message.NewMessage(uuid.New().String(), payload)
}

func NewUserLeftEvent(userID, chatRoomID uuid.UUID) *message.Message {
	payload := UserLeftPayload{
		UserID:     userID,
		ChatRoomID: chatRoomID,
	}
	return message.NewMessage(uuid.New().String(), payload)
}