package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"./config"
	"./routes"
)

func main() {
	router := gin.Default()

	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	messageRouter, err := setupMessageRouter()
	if err != nil {
		log.Fatalf("Failed to setup message router: %v", err)
	}

	routes.SetupRoutes(router, db, messageRouter)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupMessageRouter() (*message.Router, error) {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))

	router, err := message.NewRouter(message.RouterConfig{}, watermill.NewStdLogger(false, false))
	if err != nil {
		return nil, fmt.Errorf("failed to create message router: %w", err)
	}

	router.AddMiddleware(middleware.Recoverer)

	config.SetupSubscribers(pubSub, router)

	return router, nil
}