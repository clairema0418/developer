package main

import (
	"log"
	"net/http"
	"websocket/server"
	"redis/redis"
)

func main() {
	go server.CreateRoom()
	go redis.ConnectToRedis()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ConnectToRoom(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}