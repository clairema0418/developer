package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

func DBReadWriteMiddleware(db *gorm.DB, pubSub *message.PubSub) gin.HandlerFunc {
	return func(c *gin.Context) {
		readDB, writeDB := db, db
		c.Set("readDB", readDB)
		c.Set("writeDB", writeDB)

		c.Next()

		correlationID := middleware.MessageCorrelationID(c.Request.Context())
		if correlationID != "" {
			pubSub.Publish(correlationID, message.NewMessage(correlationID, nil))
		}
	}
}