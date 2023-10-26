package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/darkcat013/pad-lab-1/Gateway/api"
	"github.com/darkcat013/pad-lab-1/Gateway/api/controllers"
	"github.com/darkcat013/pad-lab-1/Gateway/cache"
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/darkcat013/pad-lab-1/Gateway/services/owner"
	"github.com/darkcat013/pad-lab-1/Gateway/services/test"
	"github.com/darkcat013/pad-lab-1/Gateway/services/veterinary"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	cacheStore := cache.GetCacheStore(cfg)
	rateLimitStore := cache.GetRateLimitStore(cfg)

	ownerService := owner.NewOwnerService(cfg)
	ownerController := controllers.NewOwnerController(ownerService, cacheStore)

	veterinaryService := veterinary.NewVeterinaryService(cfg)
	veterinaryController := controllers.NewVeterinaryController(veterinaryService, cacheStore)

	testService := test.NewTestService(cfg)
	testController := controllers.NewTestController(testService, cacheStore)

	statusController := controllers.NewStatusController()

	srv := api.NewServer(cfg, ownerController, veterinaryController, testController, statusController, rateLimitStore)

	go func() {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
		log.Info().Msg("All server connections are closed")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSEGV)

	<-quit
	log.Warn().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited properly")
}
