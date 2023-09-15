package services

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/jinzhu/gorm"

	"repositories"
	"messages"
	"handlers"
)

type DBReaderService struct {
	AccountRepo repositories.AccountRepository
	Subscriber  message.Subscriber
}

func NewDBReaderService(db *gorm.DB, subscriber message.Subscriber) *DBReaderService {
	accountRepo := repositories.NewAccountRepository(db)
	return &DBReaderService{
		AccountRepo: accountRepo,
		Subscriber:  subscriber,
	}
}

func (s *DBReaderService) Start() error {
	accountCreatedHandler := handlers.NewAccountCreatedHandler(s.AccountRepo)
	accountUpdatedHandler := handlers.NewAccountUpdatedHandler(s.AccountRepo)
	accountDeletedHandler := handlers.NewAccountDeletedHandler(s.AccountRepo)

	messageHandlerMiddleware := middleware.Retry{}
	messageHandlerMiddleware.MaxRetries = 3

	subscribeAccountCreated, err := s.Subscriber.Subscribe(messages.AccountCreated)
	if err != nil {
		return err
	}
	go processMessages(subscribeAccountCreated, accountCreatedHandler, messageHandlerMiddleware)

	subscribeAccountUpdated, err := s.Subscriber.Subscribe(messages.AccountUpdated)
	if err != nil {
		return err
	}
	go processMessages(subscribeAccountUpdated, accountUpdatedHandler, messageHandlerMiddleware)

	subscribeAccountDeleted, err := s.Subscriber.Subscribe(messages.AccountDeleted)
	if err != nil {
		return err
	}
	go processMessages(subscribeAccountDeleted, accountDeletedHandler, messageHandlerMiddleware)

	return nil
}

func processMessages(messagesChannel <-chan *message.Message, handler message.NoPublishHandlerFunc, middleware middleware.Retry) {
	for msg := range messagesChannel {
		middleware.Middleware(handler)(msg)
	}
}