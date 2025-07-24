package handlers

import (
	"encoding/json"
	"net/http"

	"chat-server/internal/client"
	"chat-server/internal/hub"
)

type Handler struct {
	hub *hub.Hub
}

func New(h *hub.Hub) *Handler {
	return &Handler{hub: h}
}

func (h *Handler) ServeWS(w http.ResponseWriter, r *http.Request) {
	client.ServeWS(h.hub, w, r)
}

func (h *Handler) ServeHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "web/index.html")
}

func (h *Handler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := h.hub.GetRoomStats()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"rooms": stats,
		"total_clients": len(h.hub.Clients),
	})
}
