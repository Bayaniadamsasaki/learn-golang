package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"performance-monitor/internal/collector"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Hub struct {
	clients   map[*websocket.Conn]bool
	collector *collector.Collector
	broadcast chan []byte
	register  chan *websocket.Conn
	unregister chan *websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*websocket.Conn]bool),
		collector:  collector.New(),
		broadcast:  make(chan []byte),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) Run() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case client := <-h.register:
				h.clients[client] = true
				log.Println("Client connected")

			case client := <-h.unregister:
				if _, ok := h.clients[client]; ok {
					delete(h.clients, client)
					client.Close()
					log.Println("Client disconnected")
				}

			case message := <-h.broadcast:
				for client := range h.clients {
					select {
					case <-time.After(1 * time.Second):
						delete(h.clients, client)
						client.Close()
					default:
						err := client.WriteMessage(websocket.TextMessage, message)
						if err != nil {
							delete(h.clients, client)
							client.Close()
						}
					}
				}

			case <-ticker.C:
				stats, err := h.collector.CollectStats()
				if err != nil {
					continue
				}

				data, err := json.Marshal(stats)
				if err != nil {
					continue
				}

				h.broadcast <- data
			}
		}
	}()
}

func (h *Hub) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	h.register <- conn

	defer func() {
		h.unregister <- conn
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}
