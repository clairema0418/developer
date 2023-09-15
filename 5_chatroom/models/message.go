package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ChatRoomID uint           `gorm:"not null" json:"chat_room_id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Content   string         `gorm:"not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}