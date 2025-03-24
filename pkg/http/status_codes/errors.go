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
