package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"performance-monitor/internal/collector"
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

	networkBytesGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_network_bytes_total",
			Help: "Network bytes sent/received",
		},
		[]string{"direction"},
	)
)

type MetricsCollector struct{}

func NewMetricsCollector() *MetricsCollector {
	prometheus.MustRegister(cpuUsageGauge)
	prometheus.MustRegister(memoryUsageGauge)
	prometheus.MustRegister(diskUsageGauge)
	prometheus.MustRegister(networkBytesGauge)
	
	return &MetricsCollector{}
}

func (mc *MetricsCollector) UpdateMetrics(stats *collector.SystemStats) {
	for i, usage := range stats.CPU {
		cpuUsageGauge.WithLabelValues(fmt.Sprintf("cpu%d", i)).Set(usage)
	}

	memoryUsageGauge.WithLabelValues("total").Set(float64(stats.Memory.Total))
	memoryUsageGauge.WithLabelValues("used").Set(float64(stats.Memory.Used))
	memoryUsageGauge.WithLabelValues("available").Set(float64(stats.Memory.Available))

	for _, disk := range stats.Disk {
		diskUsageGauge.WithLabelValues(disk.Device, disk.Mountpoint).Set(disk.UsedPercent)
	}

	networkBytesGauge.WithLabelValues("sent").Set(float64(stats.Network.BytesSent))
	networkBytesGauge.WithLabelValues("received").Set(float64(stats.Network.BytesRecv))
}
