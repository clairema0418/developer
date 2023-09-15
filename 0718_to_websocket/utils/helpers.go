package utils

import (
	"encoding/json"
	"net/http"
	"websocket/room"
)

// Helper function to handle errors
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Helper function to broadcast a message to all clients in a room
func BroadcastMessage(room *room.Room, message []byte) {
	for client := range room.Clients {
		err := client.Conn.WriteMessage(1, message)
		HandleError(err)
	}
}

// Helper function to parse JSON from http request
func ParseJSONRequest(r *http.Request, payload interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		return err
	}
	return nil
}