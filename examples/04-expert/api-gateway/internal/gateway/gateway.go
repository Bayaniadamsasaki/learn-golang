package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"api-gateway/internal/services"
	"api-gateway/pkg/config"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Gateway struct {
	config    *config.Config
	registry  *services.ServiceRegistry
	logger    *zap.Logger
	router    *mux.Router
}

func New(cfg *config.Config, logger *zap.Logger) *Gateway {
	registry := services.NewServiceRegistry(logger)
	
	for _, svc := range cfg.Services {
		for _, upstream := range svc.Upstreams {
			registry.RegisterService(svc.Name, upstream)
		}
	}

	return &Gateway{
		config:   cfg,
		registry: registry,
		logger:   logger,
		router:   mux.NewRouter(),
	}
}

func (g *Gateway) Start() error {
	g.registry.StartHealthCheck()
	g.setupRoutes()

	server := &http.Server{
		Addr:         ":" + g.config.Server.Port,
		Handler:      g.router,
		ReadTimeout:  g.config.Server.ReadTimeout,
		WriteTimeout: g.config.Server.WriteTimeout,
	}

	g.logger.Info("API Gateway starting", zap.String("port", g.config.Server.Port))
	return server.ListenAndServe()
}

func (g *Gateway) setupRoutes() {
	g.router.HandleFunc("/health", g.healthHandler).Methods("GET")
	g.router.HandleFunc("/metrics", g.metricsHandler).Methods("GET")
	g.router.PathPrefix("/api/").HandlerFunc(g.proxyHandler)
}

func (g *Gateway) proxyHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/")
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	serviceName := parts[0]
	service := g.registry.GetHealthyService(serviceName)
	if service == nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(service.URL)
	
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

	allServices := g.registry.GetAllServices()
	for serviceName, services := range allServices {
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func (g *Gateway) metricsHandler(w http.ResponseWriter, r *http.Request) {
	metrics := map[string]interface{}{
		"timestamp": time.Now(),
		"gateway": map[string]interface{}{
			"uptime": time.Since(time.Now()),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}
