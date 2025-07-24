package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Type      string    `json:"type"`
	Room      string    `json:"room"`
	Username  string    `json:"username"`
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Client struct {
	conn     *websocket.Conn
	username string
	room     string
	send     chan Message
}

type Room struct {
	name       string
	clients    map[*Client]bool
	broadcast  chan Message
	register   chan *Client
	unregister chan *Client
}

type Hub struct {
	rooms map[string]*Room
}

func main() {
	hub := &Hub{
		rooms: make(map[string]*Room),
	}
	
	go hub.run()
	
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWebSocket(hub, w, r)
	})
	
	fmt.Println("Chat server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (h *Hub) run() {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func (h *Hub) getOrCreateRoom(roomName string) *Room {
	if room, exists := h.rooms[roomName]; exists {
		return room
	}
	
	room := &Room{
		name:       roomName,
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message, 256),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	
	h.rooms[roomName] = room
	go room.run()
	
	return room
}

func (r *Room) run() {
	for {
		select {
		case client := <-r.register:
			r.clients[client] = true
			
			joinMsg := Message{
				Type:      "join",
				Room:      r.name,
				Username:  client.username,
				Content:   fmt.Sprintf("%s joined the room", client.username),
				Timestamp: time.Now(),
			}
			
			for c := range r.clients {
				select {
				case c.send <- joinMsg:
				default:
					close(c.send)
					delete(r.clients, c)
				}
			}
			
		case client := <-r.unregister:
			if _, ok := r.clients[client]; ok {
				delete(r.clients, client)
				close(client.send)
				
				leaveMsg := Message{
					Type:      "leave",
					Room:      r.name,
					Username:  client.username,
					Content:   fmt.Sprintf("%s left the room", client.username),
					Timestamp: time.Now(),
				}
				
				for c := range r.clients {
					select {
					case c.send <- leaveMsg:
					default:
						close(c.send)
						delete(r.clients, c)
					}
				}
			}
			
		case message := <-r.broadcast:
			for client := range r.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(r.clients, client)
				}
			}
		}
	}
}

func handleWebSocket(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}
	
	client := &Client{
		conn: conn,
		send: make(chan Message, 256),
	}
	
	go client.writePump()
	go client.readPump(hub)
}

func (c *Client) readPump(hub *Hub) {
	defer func() {
		if c.room != "" {
			room := hub.getOrCreateRoom(c.room)
			room.unregister <- c
		}
		c.conn.Close()
	}()
	
	c.conn.SetReadLimit(512)
	c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	
	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		
		msg.Timestamp = time.Now()
		
		switch msg.Type {
		case "join":
			c.username = msg.Username
			c.room = msg.Room
			room := hub.getOrCreateRoom(msg.Room)
			room.register <- c
			
		case "message":
			if c.room != "" {
				msg.Username = c.username
				room := hub.getOrCreateRoom(c.room)
				room.broadcast <- msg
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			
			if err := c.conn.WriteJSON(message); err != nil {
				return
			}
			
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	html := `<!DOCTYPE html>
<html>
<head>
    <title>Chat Server</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        #messages { height: 300px; overflow-y: scroll; border: 1px solid #ccc; padding: 10px; margin: 10px 0; }
        .message { margin: 5px 0; }
        .join { color: green; }
        .leave { color: red; }
        input, button { padding: 8px; margin: 5px; }
    </style>
</head>
<body>
    <h1>Chat Room</h1>
    <div>
        <input type="text" id="username" placeholder="Username" />
        <input type="text" id="room" placeholder="Room" value="general" />
        <button onclick="connect()">Join</button>
        <button onclick="disconnect()">Leave</button>
    </div>
    <div id="messages"></div>
    <div>
        <input type="text" id="messageInput" placeholder="Type a message..." onkeypress="handleKeyPress(event)" />
        <button onclick="sendMessage()">Send</button>
    </div>

    <script>
        let ws = null;
        
        function connect() {
            const username = document.getElementById('username').value;
            const room = document.getElementById('room').value;
            
            if (!username || !room) {
                alert('Please enter username and room');
                return;
            }
            
            ws = new WebSocket('ws://localhost:8080/ws');
            
            ws.onopen = function() {
                ws.send(JSON.stringify({
                    type: 'join',
                    username: username,
                    room: room
                }));
            };
            
            ws.onmessage = function(event) {
                const message = JSON.parse(event.data);
                displayMessage(message);
            };
            
            ws.onclose = function() {
                displayMessage({type: 'system', content: 'Disconnected from server'});
            };
        }
        
        function disconnect() {
            if (ws) {
                ws.close();
                ws = null;
            }
        }
        
        function sendMessage() {
            const input = document.getElementById('messageInput');
            const content = input.value.trim();
            
            if (content && ws) {
                ws.send(JSON.stringify({
                    type: 'message',
                    content: content
                }));
                input.value = '';
            }
        }
        
        function displayMessage(message) {
            const messages = document.getElementById('messages');
            const div = document.createElement('div');
            div.className = 'message ' + message.type;
            
            const time = new Date(message.timestamp).toLocaleTimeString();
            div.innerHTML = '<strong>[' + time + '] ' + (message.username || 'System') + ':</strong> ' + message.content;
            
            messages.appendChild(div);
            messages.scrollTop = messages.scrollHeight;
        }
        
        function handleKeyPress(event) {
            if (event.key === 'Enter') {
                sendMessage();
            }
        }
    </script>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}
