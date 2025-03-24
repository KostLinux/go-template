package main

import (
	"fmt"
	"go-template/config"
	"go-template/controller"
	"go-template/pkg/logger"
	"go-template/pkg/server"
	"go-template/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
		os.Exit(1)
	}

	// Initialize logger
	if err := logger.Setup(cfg); err != nil {
		log.Printf("Error initializing logger: %v", err)
		os.Exit(1)
	}

	// Set Gin mode based on environment
	if cfg.App.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize controllers
	controllers := controller.NewControllers()

	// Initialize router with config and controllers
	router := router.New(cfg, controllers)

	// Create server
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.App.Port),
		Handler:           router,
		ReadTimeout:       time.Duration(cfg.HTTP.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(cfg.HTTP.WriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(cfg.HTTP.IdleTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(cfg.HTTP.ReadTimeout) * time.Second,
		MaxHeaderBytes:    cfg.HTTP.MaxHeaderBytes << 20,
	}

	// Start server in goroutine
	go func() {
		logger.Infof("Starting server on %s in %s mode", srv.Addr, cfg.App.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorf("Server failed to start: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal and gracefully shutdown
	server.GracefulShutdown(srv)
}
