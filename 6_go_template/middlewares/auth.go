package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupMiddleware() *gin.Engine {
	r := gin.Default()
	r.Use(AuthMiddleware())
	return r
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Add your authentication logic here
		// If authentication fails, you can use c.AbortWithStatus(http.StatusUnauthorized) to stop the request

		c.Next()
	}
}