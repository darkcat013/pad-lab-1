package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type StatusController interface {
	Status(ctx *gin.Context)
}

type statusController struct {
}

func NewStatusController() StatusController {
	log.Info().Msg("Creating new status controller")

	return &statusController{}
}

func (c *statusController) Status(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
