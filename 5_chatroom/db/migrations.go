package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/yourprojectname/models"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.ChatRoom{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Message{})
}