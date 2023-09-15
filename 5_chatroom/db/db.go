package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/yourprojectname/models"
)

var DB *gorm.DB

func Connect() {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	var err error
	DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.ChatRoom{}, &models.User{}, &models.Message{})
}

func Disconnect() {
	DB.Close()
}