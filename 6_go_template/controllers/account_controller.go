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

func NewAccountController(accountService services.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
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

func (ac *AccountController) UpdateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ac.accountService.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ac *AccountController) DeleteAccount(c *gin.Context) {
	accountID := c.Param("accountID")

	err := ac.accountService.DeleteAccount(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

func (ac *AccountController) GetAccount(c *gin.Context) {
	accountID := c.Param("accountID")

	account, err := ac.accountService.GetAccount(accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (ac *AccountController) ListAccounts(c *gin.Context) {
	accounts, err := ac.accountService.ListAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, accounts)
}