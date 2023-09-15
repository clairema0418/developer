package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/services"
)

func RegisterYouTubeController(router *gin.RouterGroup) {
	router.POST("/youtube", createYouTubeVideo)
}

func createYouTubeVideo(c *gin.Context) {
	youtubeURL := c.PostForm("url")
	if youtubeURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL is required"})
		return
	}

	youtubeService := services.NewYouTubeService()
	video, err := youtubeService.CreateYouTubeVideo(youtubeURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"video": video})
}