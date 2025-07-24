# API Gateway

API Gateway microservice dengan rate limiting, load balancing, dan monitoring.

## Fitur

- **Rate Limiting**: Membatasi request per IP
- **Load Balancing**: Distribusi traffic ke backend services
- **Health Check**: Monitor status backend services
- **Metrics**: Prometheus metrics untuk monitoring
- **Logging**: Structured logging dengan Zap
- **CORS**: Cross-origin request handling

## Endpoint

- `GET /health` - Health check gateway
- `POST /api/users/*` - Route ke user service
- `POST /api/orders/*` - Route ke order service
- `GET /metrics` - Prometheus metrics

## Backend Services

Gateway akan forward request ke:

- User Service: `http://localhost:8081`
- Order Service: `http://localhost:8082`

## Running

```bash
go run main.go
```

Server berjalan di port 8080.
