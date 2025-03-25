package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-template/config"
	"go-template/controller"
	"go-template/pkg/database"
	"go-template/pkg/logger"
	"go-template/pkg/telemetry"
	"go-template/repository"
	"go-template/router"
	"go-template/service"

	_ "go-template/docs" // This is where Swag will generate docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.opentelemetry.io/otel"
)

//	@title			Go Template API
//	@version		1.0.0
//	@description	A RESTful API for different apps
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1

func main() {
	if err := bootApp(); err != nil {
		log.Printf("Application failed to start: %v", err)
		os.Exit(1)
	}
}

func bootApp() error {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Initialize logger first for proper error reporting
	if err := logger.Setup(cfg); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	// Initialize OpenTelemetry
	var cleanup func()
	if cfg.Monitoring.Telemetry.Enabled {
		cleanup = telemetry.InitTracer(cfg)
		defer cleanup()

		logger.Infof("OpenTelemetry initialized successfully")
		logger.Infof("Verifying connection with observability provider")
		tracer := otel.Tracer("main")
		telemetry.VerifyConnection(context.Background(), tracer)
	}

	// Set Gin mode based on environment
	if cfg.App.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	db, err := database.New(cfg.Database)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	if err := db.Connect(); err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(context.Background()); err != nil {
		return fmt.Errorf("database health check failed: %w", err)
	}

	logger.Infof("Database connection established successfully")

	// Initialize all layers
	repos := repository.NewRepositories(db.GetDB())
	services := service.NewServices(repos)
	controllers := controller.NewControllers(services)

	// Initialize router with config and controllers
	router := router.New(cfg, controllers)
	router.Static("/docs", "./docs")

	// Add swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/docs", func(ctx *gin.Context) {
		ctx.File("/docs/index.html")
	})

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
	gracefulShutdown(srv)
	return nil
}

func gracefulShutdown(srv *http.Server) {
	// Create channel for shutdown signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	<-quit
	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
