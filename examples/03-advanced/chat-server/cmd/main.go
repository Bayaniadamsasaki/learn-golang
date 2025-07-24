package main

import (
	"log"
	"net/http"

	"chat-server/internal/handlers"
	"chat-server/internal/hub"
)

func main() {
	h := hub.New()
	go h.Run()

	handler := handlers.New(h)

	http.HandleFunc("/", handler.ServeHome)
	http.HandleFunc("/ws", handler.ServeWS)
	http.HandleFunc("/api/stats", handler.GetStats)

	log.Println("Chat server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
