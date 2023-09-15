package services

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"chatroom/models"
	"chatroom/repositories"
	"chatroom/publishers"
)

type MessageService struct {
	db             *gorm.DB
	messageRepo    repositories.MessageRepository
	messagePublisher publishers.MessagePublisher
}

func NewMessageService(db *gorm.DB, messageRepo repositories.MessageRepository, messagePublisher publishers.MessagePublisher) *MessageService {
	return &MessageService{
		db:             db,
		messageRepo:    messageRepo,
		messagePublisher: messagePublisher,
	}
}

func (ms *MessageService) SendMessage(c *gin.Context, chatRoomID uint, userID uint, content string) (*models.Message, error) {
	message := &models.Message{
		ChatRoomID: chatRoomID,
		UserID:     userID,
		Content:    content,
	}

	err := ms.messageRepo.CreateMessage(message)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage("message_sent", message)
	err = ms.messagePublisher.PublishMessage(msg)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (ms *MessageService) GetMessagesByChatRoomID(chatRoomID uint) ([]*models.Message, error) {
	messages, err := ms.messageRepo.GetMessagesByChatRoomID(chatRoomID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}