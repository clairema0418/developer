package services

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"yourappname/messages"
	"yourappname/repositories"
)

type DBWriterService struct {
	publisher *gochannel.Publisher
	repo      repositories.DBWriterRepository
}

func NewDBWriterService(publisher *gochannel.Publisher, repo repositories.DBWriterRepository) *DBWriterService {
	return &DBWriterService{
		publisher: publisher,
		repo:      repo,
	}
}

func (s *DBWriterService) CreateAccount(account *repositories.Account) error {
	err := s.repo.CreateAccount(account)
	if err != nil {
		return err
	}

	msg := messages.NewAccountCreatedMessage(account.ID)
	err = s.publisher.Publish(messages.AccountCreatedTopic, msg)
	return err
}

func (s *DBWriterService) UpdateAccount(account *repositories.Account) error {
	err := s.repo.UpdateAccount(account)
	if err != nil {
		return err
	}

	msg := messages.NewAccountUpdatedMessage(account.ID)
	err = s.publisher.Publish(messages.AccountUpdatedTopic, msg)
	return err
}

func (s *DBWriterService) DeleteAccount(accountID uint) error {
	err := s.repo.DeleteAccount(accountID)
	if err != nil {
		return err
	}

	msg := messages.NewAccountDeletedMessage(accountID)
	err = s.publisher.Publish(messages.AccountDeletedTopic, msg)
	return err
}