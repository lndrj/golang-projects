package handlers

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
	send chan string
}

//________________________________

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var Clients = make(map[*Client]bool)
var ClientsMu sync.Mutex

var Broadcast = make(chan string)

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()

	client := &Client{conn: ws, send: make(chan string)}
	ClientsMu.Lock()
	Clients[client] = true
	ClientsMu.Unlock()

	go func() {
		for msg := range client.send {
			ws.WriteMessage(websocket.TextMessage, []byte(msg))
		}
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		Broadcast <- string(msg)
	}

	ClientsMu.Lock()
	delete(Clients, client)
	ClientsMu.Unlock()
}

func HandleMessages() {
	for {
		msg := <-Broadcast
		ClientsMu.Lock()
		for client := range Clients {
			client.send <- msg
		}
		ClientsMu.Unlock()
	}
}
