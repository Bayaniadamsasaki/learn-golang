# Performance Monitor

Real-time system performance monitoring dengan WebSocket streaming.

## Fitur

- **CPU Monitoring**: Real-time CPU usage per core
- **Memory Monitoring**: RAM dan Swap usage
- **Disk Monitoring**: Disk space dan I/O statistics
- **Network Monitoring**: Network interface statistics
- **Process Monitoring**: Top processes by CPU/Memory
- **WebSocket Streaming**: Real-time data streaming
- **Metrics Export**: Prometheus metrics endpoint

## Endpoint

- `GET /` - Dashboard web interface
- `GET /ws` - WebSocket connection untuk real-time data
- `GET /metrics` - Prometheus metrics
- `GET /api/stats` - JSON API untuk system stats

## Metrics Collected

- CPU usage percentage per core
- Memory usage (total, available, used)
- Disk usage dan read/write operations
- Network bytes sent/received
- Top 10 processes by resource usage

## Running

```bash
go run main.go
```

Dashboard tersedia di `http://localhost:8080`
