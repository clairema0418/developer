package handlers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"yourappname/messages"
	"yourappname/repositories"
)

type AccountDeletedHandler struct {
	accountRepository repositories.AccountRepository
}

func NewAccountDeletedHandler(accountRepository repositories.AccountRepository) *AccountDeletedHandler {
	return &AccountDeletedHandler{
		accountRepository: accountRepository,
	}
}

func (h *AccountDeletedHandler) Handle(msg *message.Message) error {
	accountDeleted := &messages.AccountDeleted{}
	if err := msg.Unmarshal(accountDeleted); err != nil {
		return err
	}

	if err := h.accountRepository.DeleteAccount(accountDeleted.AccountID); err != nil {
		return err
	}

	return nil
}