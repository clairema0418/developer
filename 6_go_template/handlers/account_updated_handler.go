package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/yourusername/yourappname/messages"
	"github.com/yourusername/yourappname/repositories"
	"github.com/yourusername/yourappname/services"
)

type AccountUpdatedHandler struct {
	accountRepository repositories.AccountRepository
	publisher         message.Publisher
}

func NewAccountUpdatedHandler(accountRepository repositories.AccountRepository, publisher message.Publisher) *AccountUpdatedHandler {
	return &AccountUpdatedHandler{
		accountRepository: accountRepository,
		publisher:         publisher,
	}
}

func (h *AccountUpdatedHandler) Handle(c *gin.Context) {
	var account services.AccountUpdateInput
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedAccount, err := h.accountRepository.UpdateAccount(account)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	accountUpdatedMessage := messages.NewAccountUpdated(updatedAccount.ID)
	err = h.publisher.Publish(messages.AccountUpdatedTopic, accountUpdatedMessage)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": updatedAccount})
}