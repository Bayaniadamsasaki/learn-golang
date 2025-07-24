package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"go.uber.org/zap"
)

var (
	cpuUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage_percent",
			Help: "CPU usage percentage per core",
		},
		[]string{"core"},
	)

	memoryUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_memory_usage_bytes",
			Help: "Memory usage in bytes",
		},
		[]string{"type"},
	)

	diskUsageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_disk_usage_percent",
			Help: "Disk usage percentage",
		},
		[]string{"device", "mountpoint"},
	)
)

type SystemStats struct {
	Timestamp time.Time     `json:"timestamp"`
	CPU       []float64     `json:"cpu"`
	Memory    MemoryStats   `json:"memory"`
	Disk      []DiskStats   `json:"disk"`
	Network   NetworkStats  `json:"network"`
	Processes []ProcessInfo `json:"processes"`
}

type MemoryStats struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskStats struct {
	Device      string  `json:"device"`
	Mountpoint  string  `json:"mountpoint"`
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"used_percent"`
}

type NetworkStats struct {
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
}

type ProcessInfo struct {
	PID         int32   `json:"pid"`
	Name        string  `json:"name"`
	CPUPercent  float64 `json:"cpu_percent"`
	MemoryMB    float32 `json:"memory_mb"`
	MemoryBytes uint64  `json:"memory_bytes"`
}

type Monitor struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
	logger   *zap.Logger
}

func NewMonitor() *Monitor {
	logger, _ := zap.NewProduction()
	
	return &Monitor{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		clients: make(map[*websocket.Conn]bool),
		logger:  logger,
	}
}

func (m *Monitor) collectStats() (*SystemStats, error) {
	stats := &SystemStats{
		Timestamp: time.Now(),
	}

	cpuPercents, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	stats.CPU = cpuPercents

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	stats.Memory = MemoryStats{
		Total:       vmStat.Total,
		Available:   vmStat.Available,
		Used:        vmStat.Used,
		UsedPercent: vmStat.UsedPercent,
	}

	diskStats, err := m.collectDiskStats()
	if err != nil {
		return nil, err
	}
	stats.Disk = diskStats

	netStats, err := m.collectNetworkStats()
	if err != nil {
		return nil, err
	}
	stats.Network = netStats

	processStats, err := m.collectProcessStats()
	if err != nil {
		return nil, err
	}
	stats.Processes = processStats

	return stats, nil
}

func (m *Monitor) collectDiskStats() ([]DiskStats, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	var diskStats []DiskStats
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}

		diskStats = append(diskStats, DiskStats{
			Device:      partition.Device,
			Mountpoint:  partition.Mountpoint,
			Total:       usage.Total,
			Used:        usage.Used,
			Free:        usage.Free,
			UsedPercent: usage.UsedPercent,
		})
	}

	return diskStats, nil
}

func (m *Monitor) collectNetworkStats() (NetworkStats, error) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		return NetworkStats{}, err
	}

	if len(netIO) == 0 {
		return NetworkStats{}, nil
	}

	return NetworkStats{
		BytesSent:   netIO[0].BytesSent,
		BytesRecv:   netIO[0].BytesRecv,
		PacketsSent: netIO[0].PacketsSent,
		PacketsRecv: netIO[0].PacketsRecv,
	}, nil
}

func (m *Monitor) collectProcessStats() ([]ProcessInfo, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}

	var processes []ProcessInfo
	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			continue
		}

		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		memInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}

		processes = append(processes, ProcessInfo{
			PID:         pid,
			Name:        name,
			CPUPercent:  cpuPercent,
			MemoryMB:    float32(memInfo.RSS) / 1024 / 1024,
			MemoryBytes: memInfo.RSS,
		})
	}

	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPUPercent > processes[j].CPUPercent
	})

	if len(processes) > 10 {
		processes = processes[:10]
	}

	return processes, nil
}

func (m *Monitor) updateMetrics(stats *SystemStats) {
	for i, usage := range stats.CPU {
		cpuUsageGauge.WithLabelValues(fmt.Sprintf("cpu%d", i)).Set(usage)
	}

	memoryUsageGauge.WithLabelValues("total").Set(float64(stats.Memory.Total))
	memoryUsageGauge.WithLabelValues("used").Set(float64(stats.Memory.Used))
	memoryUsageGauge.WithLabelValues("available").Set(float64(stats.Memory.Available))

	for _, disk := range stats.Disk {
		diskUsageGauge.WithLabelValues(disk.Device, disk.Mountpoint).Set(disk.UsedPercent)
	}
}

func (m *Monitor) broadcastStats() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats, err := m.collectStats()
		if err != nil {
			m.logger.Error("Failed to collect stats", zap.Error(err))
			continue
		}

		m.updateMetrics(stats)

		data, err := json.Marshal(stats)
		if err != nil {
			m.logger.Error("Failed to marshal stats", zap.Error(err))
			continue
		}

		for client := range m.clients {
			err := client.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				m.logger.Error("Failed to write message", zap.Error(err))
				client.Close()
				delete(m.clients, client)
			}
		}
	}
}

func (m *Monitor) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		m.logger.Error("Failed to upgrade connection", zap.Error(err))
		return
	}
	defer conn.Close()

	m.clients[conn] = true
	m.logger.Info("New WebSocket client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			m.logger.Info("WebSocket client disconnected")
			delete(m.clients, conn)
			break
		}
	}
}

func (m *Monitor) handleStats(w http.ResponseWriter, r *http.Request) {
	stats, err := m.collectStats()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (m *Monitor) handleDashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>Performance Monitor</title>
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
</head>
<body>
    <h1>System Performance Monitor</h1>
    <div id="stats"></div>
    <canvas id="cpuChart" width="400" height="200"></canvas>
    <canvas id="memoryChart" width="400" height="200"></canvas>
    
    <script>
        const ws = new WebSocket('ws://localhost:8080/ws');
        const statsDiv = document.getElementById('stats');
        
        ws.onmessage = function(event) {
            const data = JSON.parse(event.data);
            statsDiv.innerHTML = '<pre>' + JSON.stringify(data, null, 2) + '</pre>';
        };
    </script>
</body>
</html>`

	t, _ := template.New("dashboard").Parse(tmpl)
	t.Execute(w, nil)
}

func init() {
	prometheus.MustRegister(cpuUsageGauge)
	prometheus.MustRegister(memoryUsageGauge)
	prometheus.MustRegister(diskUsageGauge)
}

func main() {
	monitor := NewMonitor()
	defer monitor.logger.Sync()

	go monitor.broadcastStats()

	http.HandleFunc("/", monitor.handleDashboard)
	http.HandleFunc("/ws", monitor.handleWebSocket)
	http.HandleFunc("/api/stats", monitor.handleStats)
	http.Handle("/metrics", promhttp.Handler())

	monitor.logger.Info("Performance Monitor starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
