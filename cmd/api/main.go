package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-template/config"
	"go-template/controller"
	"go-template/pkg/database"
	"go-template/pkg/logger"
	"go-template/pkg/server"
	"go-template/repository"
	"go-template/router"
	"go-template/service"

	_ "go-template/docs" // This is where Swag will generate docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

//	@host		localhost:8080
//	@BasePath	/api/v1

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

	// Initialize database
	db, err := database.NewDBManager(&cfg.Database)
	if err != nil {
		logger.Errorf("Failed to create database manager: %v", err)
		os.Exit(1)
	}

	if err := db.Connect(); err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(context.Background()); err != nil {
		logger.Errorf("Database health check failed: %v", err)
		os.Exit(1)
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
	server.GracefulShutdown(srv)
}
