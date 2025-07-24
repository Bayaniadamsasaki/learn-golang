package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Server    ServerConfig    `json:"server"`
	Redis     RedisConfig     `json:"redis"`
	Services  []ServiceConfig `json:"services"`
	RateLimit RateLimitConfig `json:"rate_limit"`
	Logging   LoggingConfig   `json:"logging"`
}

type ServerConfig struct {
	Port         string        `json:"port"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
	CORS         CORSConfig    `json:"cors"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type ServiceConfig struct {
	Name      string   `json:"name"`
	Prefix    string   `json:"prefix"`
	Upstreams []string `json:"upstreams"`
	Strategy  string   `json:"strategy"`
}

type RateLimitConfig struct {
	Limit  int           `json:"limit"`
	Window time.Duration `json:"window"`
}

type LoggingConfig struct {
	Level  string `json:"level"`
	Format string `json:"format"`
}

type CORSConfig struct {
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func Default() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         "8080",
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
			CORS: CORSConfig{
				AllowedOrigins: []string{"*"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders: []string{"*"},
			},
		},
		Redis: RedisConfig{
			Addr: "localhost:6379",
			DB:   0,
		},
		Services: []ServiceConfig{
			{
				Name:      "users",
				Prefix:    "/api/users",
				Upstreams: []string{"http://localhost:8081"},
				Strategy:  "round_robin",
			},
			{
				Name:      "orders",
				Prefix:    "/api/orders",
				Upstreams: []string{"http://localhost:8082"},
				Strategy:  "round_robin",
			},
		},
		RateLimit: RateLimitConfig{
			Limit:  100,
			Window: time.Minute,
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
		},
	}
}
