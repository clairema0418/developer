```go
package account_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"app/account"
)

func TestAccountRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	account.Routes(router)

	t.Run("GET /accounts - success", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/accounts", nil)
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("POST /accounts - success", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/accounts", nil)
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})

	t.Run("PUT /accounts/:id - success", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("PUT", "/accounts/1", nil)
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("DELETE /accounts/:id - success", func(t *testing.T) {
		response := httptest.NewRecorder()
		request, _ := http.NewRequest("DELETE", "/accounts/1", nil)
		router.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}
```