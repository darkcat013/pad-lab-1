package controllers

import (
	"io"
	"net/http"

	"github.com/darkcat013/pad-lab-1/Gateway/services/veterinary"
	"github.com/darkcat013/pad-lab-1/Gateway/services/veterinary/pb"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/encoding/protojson"
)

type VeterinaryController interface {
	MakeAppointment(ctx *gin.Context)
	EndAppointment(ctx *gin.Context)
}

type veterinaryController struct {
	service veterinary.VeterinaryService
}

func NewVeterinaryController(service veterinary.VeterinaryService) VeterinaryController {
	log.Info().Msg("Creating new veterinary controller")

	return &veterinaryController{service}
}

func (c *veterinaryController) MakeAppointment(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var veterinary pb.MakeAppointmentRequest
	err = protojson.Unmarshal(bodyBytes, &veterinary)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := c.service.SendMakeAppointmentRequest(&veterinary)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": responseMessage})
}

func (c *veterinaryController) EndAppointment(ctx *gin.Context) {
	var pet pb.EndAppointmentRequest
	err := ctx.BindJSON(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := c.service.SendEndAppointmentRequest(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": responseMessage})
}
