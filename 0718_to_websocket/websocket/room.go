package websocket

import (
	"sync"
)

type Room struct {
	Name    string
	Clients map[*Client]bool
	Broadcast  chan Message
}

var rooms = make(map[string]*Room)
var mutex = &sync.Mutex{}

func createRoom(name string) *Room {
	mutex.Lock()
	defer mutex.Unlock()

	if room, ok := rooms[name]; ok {
		return room
	}

	room := &Room{
		Name:    name,
		Clients: make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}

	rooms[name] = room

	go room.run()

	return room
}

func (room *Room) run() {
	for {
		msg := <-room.Broadcast
		for client := range room.Clients {
			client.send <- msg
		}
	}
}