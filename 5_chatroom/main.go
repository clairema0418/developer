package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"os"

	"./config"
	"./db"
	"./middlewares"
	"./routes"
)

var (
	router *gin.Engine
	pub    *watermill.PubSub
)

func main() {
	setupDatabase()
	setupRouter()
	setupMiddleware()
	setupPubSub()

	router.Run(":8080")
}

func setupDatabase() {
	var err error
	db.DB, err = gorm.Open("postgres", config.GetDatabaseURL())
	if err != nil {
		panic("failed to connect to database")
	}
	db.RunMigrations()
}

func setupRouter() {
	router = gin.Default()
	routes.SetupRoutes(router)
}

func setupMiddleware() {
	router.Use(middlewares.AuthMiddleware())
}

func setupPubSub() {
	pub = watermill.NewPubSub(gochannel.NewGoChannel(gochannel.Config{}), message.NewNoopMarshaler())
}