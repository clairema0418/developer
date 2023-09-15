package subscribers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"chatroom_backend/events"
	"chatroom_backend/services"
)

type UserSubscriber struct {
	DB          *gorm.DB
	UserService *services.UserService
}

func NewUserSubscriber(db *gorm.DB, userService *services.UserService) *UserSubscriber {
	return &UserSubscriber{
		DB:          db,
		UserService: userService,
	}
}

func (us *UserSubscriber) SubscribeUserEvents(subscriber message.Subscriber) error {
	userJoined, err := subscriber.Subscribe(events.UserJoined)
	if err != nil {
		return err
	}

	userLeft, err := subscriber.Subscribe(events.UserLeft)
	if err != nil {
		return err
	}

	go us.handleUserJoined(userJoined)
	go us.handleUserLeft(userLeft)

	return nil
}

func (us *UserSubscriber) handleUserJoined(messages <-chan *message.Message) {
	for msg := range messages {
		var userJoinedEvent events.UserJoinedEvent
		err := msg.Unmarshal(&userJoinedEvent)
		if err != nil {
			continue
		}

		err = us.UserService.AddUserToChatRoom(userJoinedEvent.UserID, userJoinedEvent.ChatRoomID)
		if err != nil {
			continue
		}

		msg.Ack()
	}
}

func (us *UserSubscriber) handleUserLeft(messages <-chan *message.Message) {
	for msg := range messages {
		var userLeftEvent events.UserLeftEvent
		err := msg.Unmarshal(&userLeftEvent)
		if err != nil {
			continue
		}

		err = us.UserService.RemoveUserFromChatRoom(userLeftEvent.UserID, userLeftEvent.ChatRoomID)
		if err != nil {
			continue
		}

		msg.Ack()
	}
}