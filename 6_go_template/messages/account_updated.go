package messages

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const AccountUpdatedTopic = "account_updated"

type AccountUpdated struct {
	AccountID string
	Name      string
	Email     string
}

func (a *AccountUpdated) Marshal() (*message.Message, error) {
	payload := []byte(a.AccountID + "," + a.Name + "," + a.Email)
	msg := message.NewMessage(message.NewUUID(), payload)
	msg.Metadata.Set("type", AccountUpdatedTopic)
	return msg, nil
}

func UnmarshalAccountUpdated(msg *message.Message) (*AccountUpdated, error) {
	payload := string(msg.Payload)
	data := strings.Split(payload, ",")
	return &AccountUpdated{
		AccountID: data[0],
		Name:      data[1],
		Email:     data[2],
	}, nil
}