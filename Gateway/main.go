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
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/darkcat013/pad-lab-1/Gateway/services/owner"
	"github.com/darkcat013/pad-lab-1/Gateway/services/veterinary"
	"github.com/rs/zerolog/log"
)

// type OwnerServer struct {
// 	owner.UnsafeOwnerServer
// }

// func (s OwnerServer) Create(ctx context.Context, req *owner.CreateRequest) (*owner.CreateResponse, error) {

// 	return &owner.CreateResponse{
// 		Response: "test",
// 	}, nil
// }

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// redis, err := redis.NewDB(cfg)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to connect to database")
	// }

	ownerService := owner.NewOwnerService(cfg)
	ownerController := controllers.NewOwnerController(ownerService)

	veterinaryService := veterinary.NewVeterinaryService(cfg)
	veterinaryController := controllers.NewVeterinaryController(veterinaryService)

	srv := api.NewServer(cfg, ownerController, veterinaryController)

	// lis, err := net.Listen("tcp", ":8089")
	// if err != nil {
	// 	log.Fatalf("cannot create listener: %s", err)
	// }

	// serverRegistrar := grpc.NewServer()
	// server := &OwnerServer{}

	// owner.RegisterOwnerServer(serverRegistrar, server)

	// err = serverRegistrar.Serve(lis)
	// if err != nil {
	// 	log.Fatalf("impossible to serve: %s", err)
	// }

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

	// if err := data.CloseDB(db); err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to close db connection")
	// }

	log.Info().Msg("Server exited properly")
}
