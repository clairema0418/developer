package repositories

import (
	"github.com/jinzhu/gorm"
	"chatroom/models"
)

type ChatRoomRepository struct {
	DB *gorm.DB
}

func NewChatRoomRepository(db *gorm.DB) *ChatRoomRepository {
	return &ChatRoomRepository{DB: db}
}

func (cr *ChatRoomRepository) CreateChatRoom(chatRoom *models.ChatRoom) error {
	return cr.DB.Create(chatRoom).Error
}

func (cr *ChatRoomRepository) GetChatRoomByID(id uint) (*models.ChatRoom, error) {
	var chatRoom models.ChatRoom
	err := cr.DB.First(&chatRoom, id).Error
	if err != nil {
		return nil, err
	}
	return &chatRoom, nil
}

func (cr *ChatRoomRepository) GetChatRooms(isPublic bool) ([]*models.ChatRoom, error) {
	var chatRooms []*models.ChatRoom
	err := cr.DB.Where("is_public = ?", isPublic).Find(&chatRooms).Error
	if err != nil {
		return nil, err
	}
	return chatRooms, nil
}

func (cr *ChatRoomRepository) UpdateChatRoom(chatRoom *models.ChatRoom) error {
	return cr.DB.Save(chatRoom).Error
}

func (cr *ChatRoomRepository) DeleteChatRoom(id uint) error {
	return cr.DB.Delete(&models.ChatRoom{}, id).Error
}