package utils

const (
	// Redis connection constants
	RedisHost     = "localhost:6379"
	RedisPassword = ""
	RedisDB       = 0

	// WebSocket constants
	WebSocketProtocol = "ws"
	WebSocketPort     = ":8080"

	// Room constants
	DefaultRoomName = "general"

	// Message types
	MessageTypeBroadcast = "broadcast"
	MessageTypeRedis     = "redis"
)