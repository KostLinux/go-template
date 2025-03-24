package main

import (
	"fmt"
	"go-template/config"
	"go-template/controller"
	"go-template/pkg/server"
	"go-template/router"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading config: %v", err)
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
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Starting server on %s in %s mode", srv.Addr, cfg.App.Env)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal and gracefully shutdown
	server.GracefulShutdown(srv)
}
