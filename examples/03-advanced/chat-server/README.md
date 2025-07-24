# Chat Server

Real-time chat server menggunakan WebSocket dengan goroutines untuk menangani multiple clients.

## Fitur

- Real-time messaging dengan WebSocket
- Multiple chat rooms
- User authentication sederhana
- Broadcast messages ke semua clients
- Private messaging
- User online/offline status
- Message history

## Setup

```bash
go mod tidy
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## Cara Menggunakan

1. Buka browser ke `http://localhost:8080`
2. Masukkan username
3. Pilih atau buat chat room
4. Mulai chatting!

## WebSocket Endpoints

- `ws://localhost:8080/ws` - WebSocket connection

## Message Format

```json
{
  "type": "message",
  "room": "general",
  "username": "john",
  "content": "Hello everyone!",
  "timestamp": "2023-01-01T00:00:00Z"
}
```

## Message Types

- `join` - User bergabung ke room
- `leave` - User meninggalkan room
- `message` - Chat message
- `private` - Private message
- `user_list` - Daftar user online

## Pembelajaran

Project ini mengajarkan:

- WebSocket implementation
- Concurrent connection handling
- Channel-based message broadcasting
- Hub pattern untuk managing connections
- Real-time communication
- JSON message protocol
- Web interface untuk testing
