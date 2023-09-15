package messages

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const AccountCreatedTopic = "account_created"

type AccountCreated struct {
	ID        string
	Email     string
	FirstName string
	LastName  string
}

func (ac *AccountCreated) Marshal() (*message.Message, error) {
	payload, err := json.Marshal(ac)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)
	msg.Metadata.Set("type", AccountCreatedTopic)

	return msg, nil
}

func UnmarshalAccountCreated(msg *message.Message) (*AccountCreated, error) {
	ac := &AccountCreated{}
	err := json.Unmarshal(msg.Payload, ac)
	if err != nil {
		return nil, err
	}

	return ac, nil
}