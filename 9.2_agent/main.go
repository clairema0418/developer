package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"log"
	"os"
)

var (
	audioSegmenter     = "audio_segmentation.go"
	textConverter      = "audio_to_text.go"
	vectorConverter    = "text_to_vector.go"
	dbConnection       = "database.go"
	websocketConnection = "websocket.go"
	pubsubConnection    = "pubsub.go"
)

type AudioSegment struct {
	ID     uint
	Offset int64
	Data   []byte
}

type TextSegment struct {
	ID   uint
	Text string
}

type VectorSegment struct {
	ID     uint
	Vector []float64
}

func main() {
	router := gin.Default()

	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=audio_segmentation sslmode=disable password=postgres")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&AudioSegment{}, &TextSegment{}, &VectorSegment{})

	pubSub := gochannel.NewGoChannel(gochannel.Config{}, watermill.NewStdLogger(false, false))
	router.Use(WatermillMiddleware(pubSub))

	router.POST("/segment_audio", segmentAudio)
	router.POST("/convert_text", convertText)
	router.POST("/convert_vector", convertVector)
	router.POST("/store_in_database", storeInDatabase)
	router.GET("/progress", setupWebSocket)

	go setupPubSub(pubSub)

	router.Run(":8080")
}

func WatermillMiddleware(pubSub message.PubSub) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("pubsub", pubSub)
		c.Next()
	}
}

func segmentAudio(c *gin.Context) {
	// Implementation in audio_segmenter.go
}

func convertText(c *gin.Context) {
	// Implementation in text_converter.go
}

func convertVector(c *gin.Context) {
	// Implementation in vector_converter.go
}

func storeInDatabase(c *gin.Context) {
	// Implementation in database.go
}

func setupWebSocket(c *gin.Context) {
	// Implementation in websocket.go
}

func setupPubSub(pubSub message.PubSub) {
	// Implementation in pubsub.go
}