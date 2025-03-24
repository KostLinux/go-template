package middleware

import (
	"net/http"

	"go-template/config"

	"github.com/gin-gonic/gin"
)

func Csrf(cfg *config.New) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodPost {
			ctx.Writer.Header().Set("X-CSRF-Token", cfg.Middleware.Csrf.Key)
		}

		ctx.Next()
	}
}
