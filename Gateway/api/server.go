package api

import (
	"net/http"

	"github.com/darkcat013/pad-lab-1/Gateway/api/controllers"
	"github.com/darkcat013/pad-lab-1/Gateway/api/middleware"
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func NewServer(cfg config.Config,
	ownerController controllers.OwnerController,
	veterinaryController controllers.VeterinaryController,
	testController controllers.TestController,
	rateLimitStore *redis.Client) *http.Server {

	log.Info().Msg("Creating new server")

	rateLimitMiddleware := middleware.GetRateLimitMiddleware(rateLimitStore)

	e := gin.Default()
	e.ForwardedByClientIP = true
	e.Use(middleware.CORS(cfg.AllowOrigin))
	e.Use(rateLimitMiddleware)
	r := e.Group("/api")

	registerOwnerRoutes(r, ownerController)
	registerVeterinaryRoutes(r, veterinaryController)
	registerTestRoutes(r, testController)

	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: e,
	}
}

func registerOwnerRoutes(router *gin.RouterGroup, c controllers.OwnerController) {
	r := router.Group("/owner")
	r.POST("/register", c.Register)
	r.POST("/register-pet", c.RegisterPet)
	r.DELETE("/remove-data/:id", c.Delete)
}

func registerVeterinaryRoutes(router *gin.RouterGroup, c controllers.VeterinaryController) {
	r := router.Group("/veterinary")
	r.POST("/make-appointment", c.MakeAppointment)
	r.POST("/end-appointment", c.EndAppointment)
}

func registerTestRoutes(router *gin.RouterGroup, c controllers.TestController) {
	r := router.Group("/test")
	r.GET("/timeout", c.TestTimeout)
	r.GET("/rate-limit", c.TestRateLimit)
	r.GET("/circuit-breaker", c.TestCircuitBreaker)
}
