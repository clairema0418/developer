package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"app/middleware"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.AuthMiddleware())

	r.GET("/testAuth", func(c *gin.Context) {
		c.String(http.StatusOK, "passed")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/testAuth", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	req.Header.Set("Authorization", "Bearer invalid_token")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)

	// Add code here to set a valid token in the Authorization header
	// and assert that the status code is http.StatusOK
}