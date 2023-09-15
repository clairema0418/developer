package subscribers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourprojectname/events"
	"github.com/yourusername/yourprojectname/services"
)

type MessageSubscriber struct {
	messageService *services.MessageService
}

func NewMessageSubscriber(messageService *services.MessageService) *MessageSubscriber {
	return &MessageSubscriber{
		messageService: messageService,
	}
}

func (ms *MessageSubscriber) Subscribe(subscriber message.Subscriber) error {
	messageSentSubscription, err := subscriber.Subscribe(events.MessageSent)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messageSentSubscription {
			var messageSentEvent events.MessageSentEvent
			err := msg.Unmarshal(&messageSentEvent)
			if err != nil {
				msg.Nack()
				continue
			}

			err = ms.handleMessageSent(&messageSentEvent)
			if err != nil {
				msg.Nack()
			} else {
				msg.Ack()
			}
		}
	}()

	return nil
}

func (ms *MessageSubscriber) handleMessageSent(event *events.MessageSentEvent) error {
	message := event.Message
	chatRoomID := event.ChatRoomID

	ctx := &gin.Context{}
	err := ms.messageService.SendMessage(ctx, chatRoomID, message)
	if err != nil {
		return err
	}

	return nil
}