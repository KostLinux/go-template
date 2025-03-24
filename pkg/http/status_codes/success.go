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

func Created(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    "CREATED",
		"message": message,
	})
}

func NoContent(ctx *gin.Context, message interface{}) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"code":    "NO_CONTENT",
		"message": message,
	})
}
