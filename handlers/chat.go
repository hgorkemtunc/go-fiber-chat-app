package handlers

import (
	"sync"
	"time"

	"github.com/gofiber/websocket/v2"
)

type Message struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type ChatHandler struct {
	mu       sync.Mutex
	messages []Message
	clients  map[*websocket.Conn]bool
}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{
		messages: []Message{},
		clients:  make(map[*websocket.Conn]bool),
	}
}

func (h *ChatHandler) HandleWebSocket(c *websocket.Conn) {
	h.mu.Lock()
	h.clients[c] = true
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.clients, c)
		h.mu.Unlock()
		c.Close()
	}()

	for {
		mt, msg, err := c.ReadMessage()
		if err != nil {
			break
		}

		h.mu.Lock()
		for client := range h.clients {
			if err := client.WriteMessage(mt, msg); err != nil {
				client.Close()
				delete(h.clients, client)
			}
		}
		h.mu.Unlock()
	}
}
