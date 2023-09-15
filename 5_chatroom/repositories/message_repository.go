package repositories

import (
	"github.com/jinzhu/gorm"
	"chatroom/models"
)

type MessageRepository interface {
	CreateMessage(message *models.Message) error
	GetMessagesByChatRoomID(chatRoomID uint) ([]models.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message *models.Message) error {
	return r.db.Create(message).Error
}

func (r *messageRepository) GetMessagesByChatRoomID(chatRoomID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Where("chat_room_id = ?", chatRoomID).Find(&messages).Error
	return messages, err
}