package router

import (
	"go-template/config"
	"go-template/controller"
	"go-template/middleware"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func New(cfg *config.New, controllers controller.Controllers) *gin.Engine {
	// Create new router instance
	router := gin.Default()
	router.Use(middleware.Logger())
	if cfg.Monitoring.Telemetry.Enabled {
		router.Use(otelgin.Middleware(cfg.App.Name))
	}

	// Apply middleware if enabled in config
	if cfg.Middleware.Cors.Enabled {
		router.Use(middleware.Cors(cfg))
	}

	if cfg.Middleware.Csrf.Enabled {
		router.Use(middleware.Csrf(cfg))
	}

	// Setup routes
	APIv1Router(router, controllers)

	return router
}
