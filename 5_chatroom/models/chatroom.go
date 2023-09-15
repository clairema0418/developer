package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ChatRoom struct {
	ID        uint `gorm:"primary_key"`
	Name      string
	IsPublic  bool
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (chatRoom *ChatRoom) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

func (chatRoom *ChatRoom) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}