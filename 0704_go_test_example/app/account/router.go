package account

import (
	"github.com/gin-gonic/gin"
)

func InitAccountRoutes(router *gin.Engine) {
	accountGroup := router.Group("/accounts")
	{
		accountGroup.POST("/", CreateAccount)
		accountGroup.GET("/:id", GetAccount)
		accountGroup.PUT("/:id", UpdateAccount)
		accountGroup.DELETE("/:id", DeleteAccount)
	}
}