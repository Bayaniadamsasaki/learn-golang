package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "api_gateway_requests_total",
			Help: "Total number of requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "api_gateway_request_duration_seconds",
			Help: "Request duration in seconds",
		},
		[]string{"method", "endpoint"},
	)
)

type Service struct {
	Name     string
	URL      *url.URL
	Healthy  bool
	LastPing time.Time
}

type Gateway struct {
	services    map[string][]*Service
	rateLimiter *RateLimiter
	logger      *zap.Logger
	redis       *redis.Client
	mu          sync.RWMutex
}

type RateLimiter struct {
	requests map[string][]time.Time
	limit    int
	window   time.Duration
	mu       sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	requests := rl.requests[ip]

	var validRequests []time.Time
	for _, req := range requests {
		if now.Sub(req) < rl.window {
			validRequests = append(validRequests, req)
		}
	}

	if len(validRequests) >= rl.limit {
		return false
	}

	validRequests = append(validRequests, now)
	rl.requests[ip] = validRequests
	return true
}

func NewGateway() *Gateway {
	logger, _ := zap.NewProduction()
	
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	gateway := &Gateway{
		services:    make(map[string][]*Service),
		rateLimiter: NewRateLimiter(100, time.Minute),
		logger:      logger,
		redis:       rdb,
	}

	gateway.registerServices()
	go gateway.healthChecker()

	return gateway
}

func (g *Gateway) registerServices() {
	userService, _ := url.Parse("http://localhost:8081")
	orderService, _ := url.Parse("http://localhost:8082")

	g.services["users"] = []*Service{
		{Name: "user-service-1", URL: userService, Healthy: true},
	}
	g.services["orders"] = []*Service{
		{Name: "order-service-1", URL: orderService, Healthy: true},
	}
}

func (g *Gateway) healthChecker() {
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		g.mu.Lock()
		for serviceName, services := range g.services {
			for _, service := range services {
				g.checkServiceHealth(serviceName, service)
			}
		}
		g.mu.Unlock()
	}
}

func (g *Gateway) checkServiceHealth(serviceName string, service *Service) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(service.URL.String() + "/health")
	
	if err != nil || resp.StatusCode != http.StatusOK {
		service.Healthy = false
		g.logger.Warn("Service unhealthy", 
			zap.String("service", serviceName),
			zap.String("url", service.URL.String()))
	} else {
		service.Healthy = true
		resp.Body.Close()
	}
	
	service.LastPing = time.Now()
}

func (g *Gateway) getHealthyService(serviceName string) *Service {
	g.mu.RLock()
	defer g.mu.RUnlock()

	services := g.services[serviceName]
	for _, service := range services {
		if service.Healthy {
			return service
		}
	}
	return nil
}

func (g *Gateway) proxyHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	
	ip := strings.Split(r.RemoteAddr, ":")[0]
	if !g.rateLimiter.Allow(ip) {
		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
		g.recordMetrics(r.Method, r.URL.Path, "429", start)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		g.recordMetrics(r.Method, r.URL.Path, "400", start)
		return
	}

	serviceName := parts[0]
	service := g.getHealthyService(serviceName)
	if service == nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		g.recordMetrics(r.Method, r.URL.Path, "503", start)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(service.URL)
	proxy.ModifyResponse = func(resp *http.Response) error {
		g.recordMetrics(r.Method, r.URL.Path, fmt.Sprintf("%d", resp.StatusCode), start)
		return nil
	}

	g.logger.Info("Proxying request",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("service", serviceName),
		zap.String("target", service.URL.String()))

	proxy.ServeHTTP(w, r)
}

func (g *Gateway) healthHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now(),
		"services":  make(map[string]interface{}),
	}

	g.mu.RLock()
	for serviceName, services := range g.services {
		serviceStatus := make([]map[string]interface{}, 0)
		for _, service := range services {
			serviceStatus = append(serviceStatus, map[string]interface{}{
				"name":      service.Name,
				"url":       service.URL.String(),
				"healthy":   service.Healthy,
				"last_ping": service.LastPing,
			})
		}
		status["services"].(map[string]interface{})[serviceName] = serviceStatus
	}
	g.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (g *Gateway) recordMetrics(method, endpoint, status string, start time.Time) {
	requestsTotal.WithLabelValues(method, endpoint, status).Inc()
	requestDuration.WithLabelValues(method, endpoint).Observe(time.Since(start).Seconds())
}

func (g *Gateway) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		next.ServeHTTP(w, r)
		
		g.logger.Info("Request processed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Duration("duration", time.Since(start)))
	})
}

func init() {
	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestDuration)
}

func main() {
	gateway := NewGateway()
	defer gateway.logger.Sync()

	router := mux.NewRouter()
	
	router.HandleFunc("/health", gateway.healthHandler).Methods("GET")
	router.PathPrefix("/api/").HandlerFunc(gateway.proxyHandler)
	router.Handle("/metrics", promhttp.Handler())

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(gateway.loggingMiddleware(router))

	gateway.logger.Info("API Gateway starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
