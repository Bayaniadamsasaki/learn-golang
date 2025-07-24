package services

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/zap"
)

type Service struct {
	Name     string
	URL      *url.URL
	Healthy  bool
	LastPing time.Time
}

type ServiceRegistry struct {
	services map[string][]*Service
	logger   *zap.Logger
	mu       sync.RWMutex
}

func NewServiceRegistry(logger *zap.Logger) *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string][]*Service),
		logger:   logger,
	}
}

func (sr *ServiceRegistry) RegisterService(name string, upstream string) error {
	u, err := url.Parse(upstream)
	if err != nil {
		return err
	}

	service := &Service{
		Name:     name,
		URL:      u,
		Healthy:  true,
		LastPing: time.Now(),
	}

	sr.mu.Lock()
	sr.services[name] = append(sr.services[name], service)
	sr.mu.Unlock()

	return nil
}

func (sr *ServiceRegistry) GetHealthyService(serviceName string) *Service {
	sr.mu.RLock()
	defer sr.mu.RUnlock()

	services := sr.services[serviceName]
	for _, service := range services {
		if service.Healthy {
			return service
		}
	}
	return nil
}

func (sr *ServiceRegistry) StartHealthCheck() {
	ticker := time.NewTicker(30 * time.Second)
	go func() {
		for range ticker.C {
			sr.mu.Lock()
			for serviceName, services := range sr.services {
				for _, service := range services {
					sr.checkServiceHealth(serviceName, service)
				}
			}
			sr.mu.Unlock()
		}
	}()
}

func (sr *ServiceRegistry) checkServiceHealth(serviceName string, service *Service) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(service.URL.String() + "/health")

	if err != nil || resp.StatusCode != http.StatusOK {
		service.Healthy = false
		sr.logger.Warn("Service unhealthy",
			zap.String("service", serviceName),
			zap.String("url", service.URL.String()))
	} else {
		service.Healthy = true
		resp.Body.Close()
	}

	service.LastPing = time.Now()
}

func (sr *ServiceRegistry) GetAllServices() map[string][]*Service {
	sr.mu.RLock()
	defer sr.mu.RUnlock()
	
	result := make(map[string][]*Service)
	for k, v := range sr.services {
		result[k] = v
	}
	return result
}
