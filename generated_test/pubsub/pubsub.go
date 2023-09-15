package pubsub

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

var (
	Publisher message.Publisher
	Subscriber message.Subscriber
)

func init() {
	logger := watermill.NewStdLogger(false, false)

	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)
	Publisher = pubsub
	Subscriber = pubsub
}

func NewCorrelationIDMiddleware() *middleware.CorrelationID {
	return middleware.NewCorrelationID(func() string {
		return watermill.NewUUID()
	})
}