package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/events"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/services"
)

type MessageController struct {
	messageService services.MessageService
}

func NewMessageController(messageService services.MessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

func (mc *MessageController) SendMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMessage, err := mc.messageService.CreateMessage(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	events.PublishMessageSent(createdMessage)
	c.JSON(http.StatusCreated, createdMessage)
}

func (mc *MessageController) GetMessagesByChatRoomID(c *gin.Context) {
	chatRoomID := c.Param("chatRoomID")

	messages, err := mc.messageService.GetMessagesByChatRoomID(chatRoomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, messages)
}