package router

import (
	"go-template/controller"

	"github.com/gin-gonic/gin"
)

func APIv1Router(router *gin.Engine, controllers controller.Controllers) {
	v1 := router.Group("/api/v1")
	{
		// Health check endpoints
		v1.GET("/ping", controllers.Status().Ping)
	}
}
