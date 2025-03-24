package httpstatus

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    "OK",
		"message": message,
	})
}
