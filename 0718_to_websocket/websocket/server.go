package websocket

import (
	"net/http"
	"github.com/gorilla/websocket"
	"../utils"
)

var (
	rooms = make(map[string]*Room)
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func createRoom(roomName string) {
	rooms[roomName] = NewRoom(roomName)
}

func connectToRoom(w http.ResponseWriter, r *http.Request) {
	roomName := r.URL.Query().Get("room")
	room, ok := rooms[roomName]
	if !ok {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	client := &Client{Conn: conn, Room: room}
	room.Register <- client
	go client.Write()
	go client.Read()
}