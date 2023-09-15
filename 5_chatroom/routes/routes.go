package routes

import (
	"github.com/gin-gonic/gin"
	"chatroom/controllers"
	"chatroom/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.AuthMiddleware())

	chatroomController := controllers.NewChatroomController()
	userController := controllers.NewUserController()
	messageController := controllers.NewMessageController()

	api := router.Group("/api")
	{
		api.POST("/chatrooms", chatroomController.CreateChatroom)
		api.POST("/chatrooms/:id/join", userController.JoinChatroom)
		api.POST("/chatrooms/:id/leave", userController.LeaveChatroom)
		api.POST("/chatrooms/:id/messages", messageController.SendMessage)
	}

	return router
}