package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/services"
	"github.com/yourusername/yourprojectname/pubsub"
)

func CreateTranscription(c *gin.Context) {
	var youtubeVideo models.YouTubeVideo
	if err := c.ShouldBindJSON(&youtubeVideo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	video, err := services.CreateYouTubeVideo(youtubeVideo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pubsub.PublishYouTubeVideoCreated(video)

	c.JSON(http.StatusCreated, gin.H{"data": video})
}

func GetTranscription(c *gin.Context) {
	videoID := c.Param("id")
	transcription, err := services.GetTranscriptionByVideoID(videoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transcription})
}