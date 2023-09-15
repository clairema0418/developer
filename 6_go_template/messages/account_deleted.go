package messages

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

const AccountDeletedTopic = "account_deleted"

type AccountDeleted struct {
	AccountID string
}

func (ad *AccountDeleted) Marshal() (*message.Message, error) {
	msg := message.NewMessage(ad.AccountID, []byte(ad.AccountID))
	msg.Metadata.Set("type", "account_deleted")
	return msg, nil
}

func UnmarshalAccountDeleted(msg *message.Message) (*AccountDeleted, error) {
	return &AccountDeleted{
		AccountID: string(msg.Payload),
	}, nil
}