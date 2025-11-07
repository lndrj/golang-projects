package main

import (
	"chatroom-ws_be/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ws", handlers.HandleConnections)

	// goroutine for broadcast to all clients
	go handlers.HandleMessages()

	log.Println("Server runs at port 8080")
	r.Run(":8080")
}
