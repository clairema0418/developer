package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"

	"api"
	"database"
	"watermill"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=questions sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	sqlDB := db.DB()
	defer sqlDB.Close()

	database.RunMigrations(db)

	router := gin.Default()

	api.SetupRouter(router, db)

	publisher, subscriber := setupWatermill(db)

	messageHandler := watermill.NewMessageHandler(subscriber, db)
	messagePublisher := watermill.NewMessagePublisher(publisher)

	router.Use(func(c *gin.Context) {
		c.Set("message_publisher", messagePublisher)
		c.Next()
	})

	go messageHandler.Run()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(fmt.Sprintf(":%s", port))
}

func setupWatermill(db *gorm.DB) (message.Publisher, message.Subscriber) {
	logger := watermill.NewStdLogger(false, false)

	pubSub := message.NewGoChannel(message.GoChannelConfig{}, logger)
	middleware.PoisonQueue(pubSub, logger)

	return pubSub, pubSub
}