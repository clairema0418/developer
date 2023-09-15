package subscribers

import (
	"context"
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"

	"github.com/yourusername/yourprojectname/config"
	"github.com/yourusername/yourprojectname/events"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/repositories"
)

type AccountSubscriber struct {
	accountRepo repositories.AccountRepository
}

func NewAccountSubscriber(accountRepo repositories.AccountRepository) *AccountSubscriber {
	return &AccountSubscriber{
		accountRepo: accountRepo,
	}
}

func (s *AccountSubscriber) Subscribe(router *message.Router) {
	router.AddMiddleware(middleware.Recoverer)

	router.AddHandler(
		"account_created",
		events.AccountCreatedTopic,
		config.MessageRouter,
		s.handleAccountCreated,
	)

	router.AddHandler(
		"account_updated",
		events.AccountUpdatedTopic,
		config.MessageRouter,
		s.handleAccountUpdated,
	)

	router.AddHandler(
		"account_deleted",
		events.AccountDeletedTopic,
		config.MessageRouter,
		s.handleAccountDeleted,
	)
}

func (s *AccountSubscriber) handleAccountCreated(msg *message.Message) ([]*message.Message, error) {
	var event events.AccountCreated
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return nil, err
	}

	account := models.Account{
		ID:        event.ID,
		Email:     event.Email,
		FirstName: event.FirstName,
		LastName:  event.LastName,
	}

	err = s.accountRepo.Create(&account)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AccountSubscriber) handleAccountUpdated(msg *message.Message) ([]*message.Message, error) {
	var event events.AccountUpdated
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return nil, err
	}

	account := models.Account{
		ID:        event.ID,
		Email:     event.Email,
		FirstName: event.FirstName,
		LastName:  event.LastName,
	}

	err = s.accountRepo.Update(&account)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *AccountSubscriber) handleAccountDeleted(msg *message.Message) ([]*message.Message, error) {
	var event events.AccountDeleted
	err := json.Unmarshal(msg.Payload, &event)
	if err != nil {
		return nil, err
	}

	err = s.accountRepo.Delete(event.ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}