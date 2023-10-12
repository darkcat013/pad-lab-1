package controllers

import (
	"net/http"

	"github.com/darkcat013/pad-lab-1/Gateway/services/owner"
	"github.com/darkcat013/pad-lab-1/Gateway/services/owner/pb"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type OwnerController interface {
	Register(ctx *gin.Context)
	RegisterPet(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ownerController struct {
	service owner.OwnerService
}

func NewOwnerController(service owner.OwnerService) OwnerController {
	log.Info().Msg("Creating new owner controller")

	return &ownerController{service}
}

func (c *ownerController) Register(ctx *gin.Context) {
	var owner pb.RegisterRequest
	err := ctx.BindJSON(&owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := c.service.SendRegisterRequest(&owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": responseMessage})
}

func (c *ownerController) RegisterPet(ctx *gin.Context) {
	var pet pb.RegisterPetRequest
	err := ctx.BindJSON(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseMessage, err := c.service.SendRegisterPetRequest(&pet)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": responseMessage})
}

func (c *ownerController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	responseMessage, err := c.service.SendDeleteRequest(&pb.DeleteRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}
