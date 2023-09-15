package publishers

import (
	"encoding/json"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/myapp/models"
)

const (
	AccountCreatedTopic = "account_created"
	AccountUpdatedTopic = "account_updated"
	AccountDeletedTopic = "account_deleted"
)

type AccountPublisher struct {
	publisher *gochannel.Publisher
}

func NewAccountPublisher(publisher *gochannel.Publisher) *AccountPublisher {
	return &AccountPublisher{
		publisher: publisher,
	}
}

func (ap *AccountPublisher) PublishAccountCreated(account *models.Account) error {
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	msg := message.NewMessage(message.NewUUID(), accountJSON)
	return ap.publisher.Publish(AccountCreatedTopic, msg)
}

func (ap *AccountPublisher) PublishAccountUpdated(account *models.Account) error {
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	msg := message.NewMessage(message.NewUUID(), accountJSON)
	return ap.publisher.Publish(AccountUpdatedTopic, msg)
}

func (ap *AccountPublisher) PublishAccountDeleted(account *models.Account) error {
	accountJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}

	msg := message.NewMessage(message.NewUUID(), accountJSON)
	return ap.publisher.Publish(AccountDeletedTopic, msg)
}