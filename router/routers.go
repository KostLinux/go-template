package router

import (
	"go-template/config"
	"go-template/controller"
	"go-template/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine      *gin.Engine
	config      *config.New
	controllers controller.Controllers
}

func New(cfg *config.New, controllers controller.Controllers) *gin.Engine {
	// Create new router instance
	router := gin.Default()

	// Apply middleware if enabled in config
	if cfg.Middleware.Cors.Enabled {
		router.Use(middleware.Cors(router, cfg))
	}

	if cfg.Middleware.Csrf.Enabled {
		router.Use(middleware.Csrf(router, cfg))
	}

	// Setup routes
	APIv1Router(router, controllers)

	return router
}
