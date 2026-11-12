package config

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients   map[*websocket.Conn]bool
	Broadcast chan []byte
	mu        sync.Mutex
}

func CreateHub() *Hub {
	return &Hub{
		clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		message := <-h.Broadcast
		h.mu.Lock()

		for client := range h.clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				client.Close()
				delete(h.clients, client)
			}
		}

		h.mu.Unlock()
	}
}

func (h *Hub) AddClient(conn *websocket.Conn) {
	h.mu.Lock()

	h.clients[conn] = true
	h.mu.Unlock()
}

func (h *Hub) RemoveClient(conn *websocket.Conn) {
	h.mu.Lock()

	delete(h.clients, conn)
	h.mu.Unlock()
}
