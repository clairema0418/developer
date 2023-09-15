package account_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"app/account"
)

func TestCreateAccount(t *testing.T) {
	router := gin.Default()
	router.POST("/accounts", account.CreateAccount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/accounts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAccount(t *testing.T) {
	router := gin.Default()
	router.GET("/accounts/:id", account.GetAccount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateAccount(t *testing.T) {
	router := gin.Default()
	router.PUT("/accounts/:id", account.UpdateAccount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/accounts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteAccount(t *testing.T) {
	router := gin.Default()
	router.DELETE("/accounts/:id", account.DeleteAccount)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/accounts/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}