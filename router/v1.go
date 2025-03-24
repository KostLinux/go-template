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

		// User endpoints
		users := v1.Group("/users")
		{
			users.POST("", controllers.User().Create)
			users.GET("", controllers.User().Get)
			users.GET("/:id", controllers.User().GetByID)
			users.PUT("/:id", controllers.User().Update)
			users.DELETE("/:id", controllers.User().Delete)
		}
	}
}
