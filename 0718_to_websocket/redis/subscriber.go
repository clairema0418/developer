package redis

import (
	"github.com/go-redis/redis"
	"../websocket"
	"../utils"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     utils.RedisAddr,
		Password: utils.RedisPassword,
		DB:       utils.RedisDB,
	})
}

func ListenToRedisTopic(topic string) {
	pubsub := redisClient.Subscribe(topic)
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	ch := pubsub.Channel()

	for msg := range ch {
		websocket.BroadcastMessage(msg.Payload)
	}
}