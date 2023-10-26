package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/darkcat013/pad-lab-1/Gateway/services/owner"
	"github.com/darkcat013/pad-lab-1/Gateway/services/owner/pb"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type OwnerController interface {
	Register(ctx *gin.Context)
	RegisterPet(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetPets(ctx *gin.Context)
}

type ownerController struct {
	service    owner.OwnerService
	cacheStore *redis.Client
}

func NewOwnerController(service owner.OwnerService, cacheStore *redis.Client) OwnerController {
	log.Info().Msg("Creating new owner controller")

	return &ownerController{service, cacheStore}
}

func (c *ownerController) Register(ctx *gin.Context) {
	var owner pb.RegisterRequest
	err := ctx.BindJSON(&owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := c.service.SendRegisterRequest(&owner)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)
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

	responseMessage, err := c.service.SendDeleteRequest(&pb.DeleteRequest{OwnerId: id, Email: id})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}

type Pet struct {
	Type string
	Race string
	Name string
}

type PetsResp struct {
	Message string
	Pets    []Pet
}

func (c *ownerController) GetPets(ctx *gin.Context) {
	id := ctx.Param("id")
	cacheId := "get-pet:" + id
	cached, err := c.cacheStore.Get(ctx, cacheId).Result()
	var response *pb.GetPetsResponse

	if cached != "" {
		errMarshal := json.Unmarshal([]byte(cached), &response)
		if errMarshal != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Info().Msg("CACHED")
		log.Info().Msg(cached)
	} else {
		response, err = c.service.SendGetPetsRequest(&pb.GetPetsRequest{OwnerId: id, Email: id})
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		jsonMarshaler := &jsonpb.Marshaler{}
		jsonStr, err := jsonMarshaler.MarshalToString(response)
		if err != nil {
			log.Err(err).Msg("Could not convert response to string")
		}
		_, err = c.cacheStore.Set(ctx, cacheId, jsonStr, 15*time.Second).Result()
		if err != nil {
			log.Err(err).Msg("Could not cache request")
		}
	}

	ctx.JSON(http.StatusOK, response)
}
