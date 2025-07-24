package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	"performance-monitor/internal/collector"
	"performance-monitor/internal/metrics"
	"performance-monitor/internal/websocket"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	hub := websocket.NewHub()
	go hub.Run()

	metricsCollector := metrics.NewMetricsCollector()
	systemCollector := collector.New()

	go func() {
		for {
			stats, err := systemCollector.CollectStats()
			if err != nil {
				continue
			}
			metricsCollector.UpdateMetrics(stats)
		}
	}()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", hub.HandleWebSocket)
	http.HandleFunc("/api/stats", handleStats(systemCollector))
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Performance Monitor starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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

	data, err := os.ReadFile("web/dashboard.html")
	if err != nil {
		http.Error(w, "Could not load dashboard", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.New("dashboard").Parse(string(data))
	if err != nil {
		http.Error(w, "Could not parse template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handleStats(collector *collector.Collector) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stats, err := collector.CollectStats()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stats)
	}
}
