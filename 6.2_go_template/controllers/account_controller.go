package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/services"
)

type AccountController struct {
	accountService services.AccountService
}

func NewAccountController(service services.AccountService) *AccountController {
	return &AccountController{accountService: service}
}

func (ac *AccountController) CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ac.accountService.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (ac *AccountController) GetAccount(c *gin.Context) {
	id := c.Param("id")

	account, err := ac.accountService.GetAccount(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) UpdateAccount(c *gin.Context) {
	id := c.Param("id")

	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAccount, err := ac.accountService.UpdateAccount(id, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAccount)
}

func (ac *AccountController) DeleteAccount(c *gin.Context) {
	id := c.Param("id")

	err := ac.accountService.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}