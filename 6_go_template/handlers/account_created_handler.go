package handlers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"yourappname/messages"
	"yourappname/repositories"
	"yourappname/services"
)

type AccountCreatedHandler struct {
	DBWriterService *services.DBWriterService
}

func NewAccountCreatedHandler(db *gorm.DB) *AccountCreatedHandler {
	accountRepository := repositories.NewAccountRepository(db)
	dbWriterService := services.NewDBWriterService(accountRepository)

	return &AccountCreatedHandler{
		DBWriterService: dbWriterService,
	}
}

func (h *AccountCreatedHandler) Handle(msg *message.Message) error {
	accountCreated := &messages.AccountCreated{}
	err := msg.Unmarshal(accountCreated)
	if err != nil {
		return err
	}

	account := &gin.Account{
		ID:        accountCreated.AccountID,
		Email:     accountCreated.Email,
		FirstName: accountCreated.FirstName,
		LastName:  accountCreated.LastName,
	}

	err = h.DBWriterService.CreateAccount(account)
	if err != nil {
		return err
	}

	return nil
}