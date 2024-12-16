// Package main adalah entry point utama aplikasi
package main

import (
	"info-retrieval/internal/app"
	"info-retrieval/internal/config"
	"info-retrieval/pkg/logger"
)

func main() {
	cfg := config.LoadConfig()
	logger.Info("Starting application on port %s", cfg.Port)
	
	if err := app.Run(cfg); err != nil {
		logger.Fatal("Failed to start application: %v", err)
	}
}