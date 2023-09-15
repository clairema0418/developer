package services

import (
	"errors"
	"github.com/yourusername/yourproject/models"
	"github.com/yourusername/yourproject/repositories"
	"github.com/yourusername/yourproject/utils"
)

type ChatRoomService interface {
	CreateChatRoom(name string, isPublic bool, password string) (*models.ChatRoom, error)
	JoinChatRoom(roomID uint, password string) (*models.ChatRoom, error)
}

type chatRoomServiceImpl struct {
	chatRoomRepo repositories.ChatRoomRepository
}

func NewChatRoomService(chatRoomRepo repositories.ChatRoomRepository) ChatRoomService {
	return &chatRoomServiceImpl{
		chatRoomRepo: chatRoomRepo,
	}
}

func (s *chatRoomServiceImpl) CreateChatRoom(name string, isPublic bool, password string) (*models.ChatRoom, error) {
	hashedPassword := ""
	if !isPublic {
		var err error
		hashedPassword, err = utils.HashPassword(password)
		if err != nil {
			return nil, err
		}
	}

	chatRoom := &models.ChatRoom{
		Name:     name,
		IsPublic: isPublic,
		Password: hashedPassword,
	}

	err := s.chatRoomRepo.Create(chatRoom)
	if err != nil {
		return nil, err
	}

	return chatRoom, nil
}

func (s *chatRoomServiceImpl) JoinChatRoom(roomID uint, password string) (*models.ChatRoom, error) {
	chatRoom, err := s.chatRoomRepo.FindByID(roomID)
	if err != nil {
		return nil, err
	}

	if !chatRoom.IsPublic {
		if !utils.CheckPasswordHash(password, chatRoom.Password) {
			return nil, errors.New("incorrect password")
		}
	}

	return chatRoom, nil
}