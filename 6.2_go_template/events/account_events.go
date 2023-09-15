package events

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	AccountCreated = "account_created"
	AccountUpdated = "account_updated"
	AccountDeleted = "account_deleted"
)

type AccountEvent struct {
	ID        uint
	Email     string
	FirstName string
	LastName  string
}

func NewAccountCreatedEvent(account *AccountEvent) *message.Message {
	msg := message.NewMessage(watermill.NewUUID(), account)
	msg.Metadata.Set("type", AccountCreated)
	return msg
}

func NewAccountUpdatedEvent(account *AccountEvent) *message.Message {
	msg := message.NewMessage(watermill.NewUUID(), account)
	msg.Metadata.Set("type", AccountUpdated)
	return msg
}

func NewAccountDeletedEvent(account *AccountEvent) *message.Message {
	msg := message.NewMessage(watermill.NewUUID(), account)
	msg.Metadata.Set("type", AccountDeleted)
	return msg
}