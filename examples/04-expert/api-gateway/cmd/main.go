package main

import (
	"flag"
	"log"

	"api-gateway/internal/gateway"
	"api-gateway/internal/middleware"
	"api-gateway/pkg/config"

	"github.com/rs/cors"
	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	var cfg *config.Config
	var err error

	if *configPath != "" {
		cfg, err = config.Load(*configPath)
		if err != nil {
			log.Fatalf("Failed to load config: %v", err)
		}
	} else {
		cfg = config.Default()
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}
	defer logger.Sync()

	gateway := gateway.New(cfg, logger)

	rateLimiter := middleware.NewRateLimiter(cfg.RateLimit.Limit, cfg.RateLimit.Window)
	loggingMiddleware := middleware.NewLoggingMiddleware(logger)
	metricsMiddleware := middleware.NewMetricsMiddleware()

	c := cors.New(cors.Options{
		AllowedOrigins: cfg.Server.CORS.AllowedOrigins,
		AllowedMethods: cfg.Server.CORS.AllowedMethods,
		AllowedHeaders: cfg.Server.CORS.AllowedHeaders,
	})

	handler := c.Handler(
		rateLimiter.Middleware(
			loggingMiddleware.Middleware(
				metricsMiddleware.Middleware(gateway.router))))

	if err := gateway.Start(); err != nil {
		logger.Fatal("Failed to start gateway", zap.Error(err))
	}
}
