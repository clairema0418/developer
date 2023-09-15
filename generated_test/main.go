package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"config"
	"controllers"
	"middlewares"
	"models"
	"pubsub"
	"utils"
	"workers"
)

func main() {
	// Initialize configuration
	config.InitConfig()

	// Initialize database connection
	db := utils.InitDB()
	defer db.Close()

	// Initialize Gin router
	router := gin.Default()

	// Initialize Watermill message broker
	pubSub := initPubSub()

	// Initialize message handlers and producers
	initMessageHandlers(pubSub)
	initMessageProducers(pubSub)

	// Initialize routes and middlewares
	middlewares.InitDBReadWriteMiddleware(db)
	controllers.InitRoutes(router)

	// Run the server
	router.Run(config.Config.Server.Address)
}

func initPubSub() *gochannel.GoChannel {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
	pubsub.InitPubSub(pubSub)
	return pubSub
}

func initMessageHandlers(pubSub *gochannel.GoChannel) {
	workers.InitYouTubeWorker(pubSub)
	workers.InitTranscriptionWorker(pubSub)
}

func initMessageProducers(pubSub *gochannel.GoChannel) {
	pubsub.InitPublishers(pubSub)
}