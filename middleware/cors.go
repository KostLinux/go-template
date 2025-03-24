package middleware

import (
	"go-template/config"
	"go-template/pkg/stringify"
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors(router *gin.Engine, cfg *config.New) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Convert slice to string with comma separation
		origins := strings.Join(cfg.Middleware.Cors.AllowOrigins, ",")
		methods := strings.Join(cfg.Middleware.Cors.AllowMethods, ",")
		headers := strings.Join(cfg.Middleware.Cors.AllowHeaders, ",")
		exposeHeaders := strings.Join(cfg.Middleware.Cors.ExposeHeaders, ",")

		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origins)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", headers)
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", stringify.BoolToString(cfg.Middleware.Cors.AllowCredentials))
		ctx.Writer.Header().Set("Access-Control-Max-Age", stringify.ToInteger(cfg.Middleware.Cors.MaxAge))

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
			return
		}

		ctx.Next()
	}
}
