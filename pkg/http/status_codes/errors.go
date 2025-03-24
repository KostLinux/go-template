package httpstatus

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, message string, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code":    "BAD_REQUEST",
		"message": message,
		"error":   err.Error(),
	})
}

func NotFound(ctx *gin.Context, message string, err error) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code":    "NOT_FOUND",
		"message": message,
		"error":   err.Error(),
	})
}

func Conflict(ctx *gin.Context, message string, err error) {
	ctx.JSON(http.StatusConflict, gin.H{
		"code":    "CONFLICT",
		"message": message,
		"error":   err.Error(),
	})
}

func InternalServerError(ctx *gin.Context, message string, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    "INTERNAL_SERVER_ERROR",
		"message": message,
		"error":   err.Error(),
	})
}
