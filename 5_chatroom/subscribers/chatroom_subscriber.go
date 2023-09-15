package subscribers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"chatroom_backend/events"
	"chatroom_backend/services"
)

type ChatRoomSubscriber struct {
	db          *gorm.DB
	chatService *services.ChatRoomService
}

func NewChatRoomSubscriber(db *gorm.DB, chatService *services.ChatRoomService) *ChatRoomSubscriber {
	return &ChatRoomSubscriber{
		db:          db,
		chatService: chatService,
	}
}

func (s *ChatRoomSubscriber) Subscribe(subscriber message.Subscriber) error {
	chatRoomCreated, err := subscriber.Subscribe(events.ChatRoomCreated)
	if err != nil {
		return err
	}

	go s.handleChatRoomCreated(chatRoomCreated)
	return nil
}

func (s *ChatRoomSubscriber) handleChatRoomCreated(messages <-chan *message.Message) {
	for msg := range messages {
		event := events.ChatRoomCreatedPayload{}
		err := msg.Unmarshal(&event)
		if err != nil {
			continue
		}

		chatRoom, err := s.chatService.CreateChatRoom(event.Name, event.IsPublic, event.Password)
		if err != nil {
			continue
		}

		msg.Ack()
	}
}