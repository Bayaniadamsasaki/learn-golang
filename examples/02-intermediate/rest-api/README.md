# REST API Server

REST API sederhana untuk manajemen user menggunakan Gorilla Mux.

## Fitur

- CRUD operations untuk User
- JSON response
- HTTP status codes
- Route handling dengan Gorilla Mux
- In-memory storage

## Setup

```bash
go mod tidy
go run main.go
```

Server akan berjalan di `http://localhost:8080`

## API Endpoints

### GET /users

Mendapatkan semua users

### GET /users/{id}

Mendapatkan user berdasarkan ID

### POST /users

Membuat user baru

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

### PUT /users/{id}

Update user berdasarkan ID

### DELETE /users/{id}

Hapus user berdasarkan ID

## Contoh Penggunaan

```bash
# Mendapatkan semua users
curl http://localhost:8080/users

# Membuat user baru
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com"}'

# Mendapatkan user berdasarkan ID
curl http://localhost:8080/users/1
```

## Pembelajaran

Project ini mengajarkan:

- HTTP server dengan net/http
- REST API design
- JSON marshaling/unmarshaling
- Route handling
- Error handling untuk HTTP
- Third-party packages (Gorilla Mux)
