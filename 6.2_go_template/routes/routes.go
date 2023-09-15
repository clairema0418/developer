package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	accountController := controllers.NewAccountController()

	api := router.Group("/api")
	{
		api.POST("/accounts", accountController.CreateAccount)
		api.GET("/accounts/:id", accountController.GetAccount)
		api.PUT("/accounts/:id", accountController.UpdateAccount)
		api.DELETE("/accounts/:id", accountController.DeleteAccount)
	}

	return router
}