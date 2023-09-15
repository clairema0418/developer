package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/yourusername/yourproject/database"
	"github.com/yourusername/yourproject/python_generator"
	"github.com/yourusername/yourproject/watermill"
)

func GenerateQuestionsHandler(c *gin.Context) {
	questionCountStr := c.Query("questionCount")
	questionCount, err := strconv.Atoi(questionCountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question count"})
		return
	}

	err = callPythonScript(questionCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call Python script"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Question generation started"})
}

func callPythonScript(questionCount int) error {
	err := python_generator.GenerateQuestions(questionCount)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), []byte(strconv.Itoa(questionCount)))
	err = watermill.MessagePublisher.Publish("questionGenerationCompleted", msg)
	if err != nil {
		return err
	}

	return nil
}

func QuestionGenerationCompletedHandler(msg *message.Message) error {
	questionCountStr := string(msg.Payload)
	questionCount, err := strconv.Atoi(questionCountStr)
	if err != nil {
		return err
	}

	questions, err := python_generator.GetGeneratedQuestions(questionCount)
	if err != nil {
		return err
	}

	err = saveQuestionsToDatabase(questions)
	if err != nil {
		return err
	}

	err = updateProgressInDB(questionCount)
	if err != nil {
		return err
	}

	return nil
}

func saveQuestionsToDatabase(questions []database.Question) error {
	for _, question := range questions {
		err := database.SaveQuestion(&question)
		if err != nil {
			return err
		}
	}
	return nil
}

func updateProgressInDB(questionCount int) error {
	err := database.UpdateProgress(questionCount)
	if err != nil {
		return err
	}
	return nil
}