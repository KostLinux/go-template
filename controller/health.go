package controller

import (
	httpstatus "go-template/pkg/http/status_codes"

	"github.com/gin-gonic/gin"
)

type StatusChecker interface {
	Ping(ctx *gin.Context)
}

type statusController struct{}

func NewStatusController() StatusChecker {
	return &statusController{}
}

func (status *statusController) Ping(ctx *gin.Context) {
	httpstatus.OK(ctx, "pong")
}
