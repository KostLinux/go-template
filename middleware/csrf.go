package middleware

import (
	"go-template/config"

	"github.com/gin-gonic/gin"
)

func Csrf(router *gin.Engine, cfg *config.New) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == "POST" {
			ctx.Writer.Header().Set("X-CSRF-Token", cfg.Middleware.Csrf.Key)
		}

		ctx.Next()
	}
}
