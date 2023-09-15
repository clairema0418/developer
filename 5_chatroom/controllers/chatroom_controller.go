package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/services"
)

func CreateChatRoom(c *gin.Context) {
	var chatRoom models.ChatRoom
	if err := c.ShouldBindJSON(&chatRoom); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdChatRoom, err := services.CreateChatRoom(&chatRoom)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdChatRoom)
}

func JoinChatRoom(c *gin.Context) {
	var joinRequest models.JoinChatRoomRequest
	if err := c.ShouldBindJSON(&joinRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetInt("userID")
	joinedChatRoom, err := services.JoinChatRoom(userID, joinRequest.RoomID, joinRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, joinedChatRoom)
}