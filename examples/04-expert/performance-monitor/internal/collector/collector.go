package collector

import (
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
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

type Collector struct{}

func New() *Collector {
	return &Collector{}
}

func (c *Collector) CollectStats() (*SystemStats, error) {
	stats := &SystemStats{
		Timestamp: time.Now(),
	}

	cpuPercents, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, err
	}
	stats.CPU = cpuPercents

	memStats, err := c.collectMemoryStats()
	if err != nil {
		return nil, err
	}
	stats.Memory = memStats

	diskStats, err := c.collectDiskStats()
	if err != nil {
		return nil, err
	}
	stats.Disk = diskStats

	netStats, err := c.collectNetworkStats()
	if err != nil {
		return nil, err
	}
	stats.Network = netStats

	processStats, err := c.collectProcessStats()
	if err != nil {
		return nil, err
	}
	stats.Processes = processStats

	return stats, nil
}

func (c *Collector) collectMemoryStats() (MemoryStats, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return MemoryStats{}, err
	}

	return MemoryStats{
		Total:       vmStat.Total,
		Available:   vmStat.Available,
		Used:        vmStat.Used,
		UsedPercent: vmStat.UsedPercent,
	}, nil
}

func (c *Collector) collectDiskStats() ([]DiskStats, error) {
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

func (c *Collector) collectNetworkStats() (NetworkStats, error) {
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

func (c *Collector) collectProcessStats() ([]ProcessInfo, error) {
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
