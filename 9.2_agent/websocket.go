package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ThreeDotsLabs/watermill/message"
	"log"
	"net/http"
)

var websocketConnection *websocket.Conn

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Progress struct {
	Percentage float64 `json:"percentage"`
	Message    string  `json:"message"`
}

func setupWebSocket(router *gin.Engine) {
	router.GET("/ws", func(c *gin.Context) {
		var err error
		websocketConnection, err = upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("Failed to set websocket upgrade: %+v", err)
			return
		}
		defer websocketConnection.Close()

		for {
			_, _, err := websocketConnection.ReadMessage()
			if err != nil {
				log.Printf("Error during websocket read: %+v", err)
				break
			}
		}
	})
}

func sendProgress(progress Progress) {
	if websocketConnection != nil {
		progressJSON, err := json.Marshal(progress)
		if err != nil {
			log.Printf("Error marshalling progress: %+v", err)
			return
		}

		err = websocketConnection.WriteMessage(websocket.TextMessage, progressJSON)
		if err != nil {
			log.Printf("Error sending progress via websocket: %+v", err)
		}
	}
}

func setupPubSub() {
	pubsubConnection, err := message.NewPubSub()
	if err != nil {
		log.Fatalf("Error setting up pubsub: %+v", err)
	}

	audioSegmented, err := pubsubConnection.Subscribe("audioSegmented")
	if err != nil {
		log.Fatalf("Error subscribing to audioSegmented: %+v", err)
	}

	textConverted, err := pubsubConnection.Subscribe("textConverted")
	if err != nil {
		log.Fatalf("Error subscribing to textConverted: %+v", err)
	}

	vectorConverted, err := pubsubConnection.Subscribe("vectorConverted")
	if err != nil {
		log.Fatalf("Error subscribing to vectorConverted: %+v", err)
	}

	go func() {
		for {
			select {
			case <-audioSegmented:
				sendProgress(Progress{Percentage: 33, Message: "Audio segmented"})
			case <-textConverted:
				sendProgress(Progress{Percentage: 66, Message: "Text converted"})
			case <-vectorConverted:
				sendProgress(Progress{Percentage: 100, Message: "Vector converted"})
			}
		}
	}()
}

func main() {
	router := gin.Default()

	setupWebSocket(router)
	setupPubSub()

	router.Run(":8080")
}