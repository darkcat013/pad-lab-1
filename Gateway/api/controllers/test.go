package controllers

import (
	"net/http"

	"github.com/darkcat013/pad-lab-1/Gateway/services/test"
	"github.com/darkcat013/pad-lab-1/Gateway/services/test/pb"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type TestController interface {
	TestTimeout(ctx *gin.Context)
	TestRateLimit(ctx *gin.Context)
	TestCircuitBreaker(ctx *gin.Context)
}

type testController struct {
	service test.TestService
}

func NewTestController(service test.TestService) TestController {
	log.Info().Msg("Creating new test controller")

	return &testController{service}
}

func (c *testController) TestTimeout(ctx *gin.Context) {
	var timeoutReq pb.TestTimeoutRequest

	responseMessage, err := c.service.TestTimeout(&timeoutReq)
	if err != nil {
		ctx.JSON(http.StatusRequestTimeout, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}

func (c *testController) TestRateLimit(ctx *gin.Context) {
	var rateLimitReq pb.TestRateLimitRequest

	responseMessage, err := c.service.TestRateLimit(&rateLimitReq)
	if err != nil {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}

func (c *testController) TestCircuitBreaker(ctx *gin.Context) {
	var circuitBreakerReq pb.TestCircuitBreakerRequest

	responseMessage, err := c.service.TestCircuitBreaker(&circuitBreakerReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": responseMessage})
}
