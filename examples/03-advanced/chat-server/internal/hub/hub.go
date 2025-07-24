package hub

import (
	"encoding/json"
	"log"
	"time"
)

type Client interface {
	GetSend() chan []byte
	GetUsername() string
	GetRoom() string
	Close()
}

type Hub struct {
	Clients    map[Client]bool
	Broadcast  chan []byte
	Register   chan Client
	Unregister chan Client
	Rooms      map[string]map[Client]bool
}

type Message struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Content  string `json:"content"`
	Room     string `json:"room"`
	Time     string `json:"time"`
}

func New() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan Client),
		Unregister: make(chan Client),
		Clients:    make(map[Client]bool),
		Rooms:      make(map[string]map[Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			room := client.GetRoom()
			
			if h.Rooms[room] == nil {
				h.Rooms[room] = make(map[Client]bool)
			}
			h.Rooms[room][client] = true

			joinMsg := Message{
				Type:     "join",
				Username: client.GetUsername(),
				Content:  client.GetUsername() + " joined the room",
				Room:     room,
				Time:     time.Now().Format("15:04:05"),
			}
			
			if msgBytes, err := json.Marshal(joinMsg); err == nil {
				h.broadcastToRoom(msgBytes, room)
			}

			log.Printf("Client %s joined room %s", client.GetUsername(), room)

		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				room := client.GetRoom()
				delete(h.Clients, client)
				delete(h.Rooms[room], client)
				
				if len(h.Rooms[room]) == 0 {
					delete(h.Rooms, room)
				}

				leaveMsg := Message{
					Type:     "leave",
					Username: client.GetUsername(),
					Content:  client.GetUsername() + " left the room",
					Room:     room,
					Time:     time.Now().Format("15:04:05"),
				}
				
				if msgBytes, err := json.Marshal(leaveMsg); err == nil {
					h.broadcastToRoom(msgBytes, room)
				}

				client.Close()
				log.Printf("Client %s left room %s", client.GetUsername(), room)
			}

		case message := <-h.Broadcast:
			var msg Message
			if err := json.Unmarshal(message, &msg); err != nil {
				continue
			}
			
			h.broadcastToRoom(message, msg.Room)
		}
	}
}

func (h *Hub) broadcastToRoom(message []byte, room string) {
	if roomClients, exists := h.Rooms[room]; exists {
		for client := range roomClients {
			select {
			case client.GetSend() <- message:
			default:
				delete(h.Clients, client)
				delete(roomClients, client)
				client.Close()
			}
		}
	}
}

func (h *Hub) GetRoomStats() map[string]int {
	stats := make(map[string]int)
	for room, clients := range h.Rooms {
		stats[room] = len(clients)
	}
	return stats
}
