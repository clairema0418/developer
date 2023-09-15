package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/controllers"
	"github.com/yourusername/yourprojectname/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.SetupMiddleware())

	accountController := controllers.NewAccountController()

	api := router.Group("/api")
	{
		account := api.Group("/account")
		{
			account.POST("/", accountController.CreateAccount)
			account.GET("/:id", accountController.GetAccount)
			account.GET("/", accountController.ListAccounts)
			account.PUT("/:id", accountController.UpdateAccount)
			account.DELETE("/:id", accountController.DeleteAccount)
		}
	}

	return router
}