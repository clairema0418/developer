package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func SetupRouter(messagePublisher *message.Publisher, messageSubscriber *message.Subscriber) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/generate_questions", func(c *gin.Context) {
			generateQuestionsHandler(c, messagePublisher)
		})
	}

	go func() {
		messageHandler(messageSubscriber)
	}()

	return router
}

func generateQuestionsHandler(c *gin.Context, messagePublisher *message.Publisher) {
	questionCount := c.Query("question_count")
	err := callPythonScript(questionCount)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to call Python script"})
		return
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte(questionCount))
	err = messagePublisher.Publish("questionGenerationCompleted", msg)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to publish message"})
		return
	}

	c.JSON(200, gin.H{"message": "Question generation started"})
}

func messageHandler(messageSubscriber *message.Subscriber) {
	messages, err := messageSubscriber.Subscribe("questionGenerationCompleted")
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		questionCount := string(msg.Payload)
		saveQuestionsToDatabase(questionCount)
		updateProgressInDB(questionCount)
		msg.Ack()
	}
}

func initWatermill() (*message.Publisher, *message.Subscriber) {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, middleware.Retry{})
	publisher := message.NewPublisher(pubSub, nil)
	subscriber := message.NewSubscriber(pubSub, nil)

	return &publisher, &subscriber
}