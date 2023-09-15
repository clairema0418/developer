package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var redisClient *redis.Client

func connectToRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}