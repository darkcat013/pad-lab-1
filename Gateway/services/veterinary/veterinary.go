package veterinary

import (
	"context"

	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/darkcat013/pad-lab-1/Gateway/services/veterinary/pb"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type VeterinaryService interface {
	SendMakeAppointmentRequest(request *pb.MakeAppointmentRequest) (string, error)
	SendEndAppointmentRequest(request *pb.EndAppointmentRequest) (string, error)
}

type veterinaryService struct {
	client pb.VeterinaryClient
}

func NewVeterinaryService(cfg config.Config) VeterinaryService {
	log.Info().Msg("Creating new veterinary service")

	conn, err := grpc.Dial(cfg.VeterinaryUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}

	client := pb.NewVeterinaryClient(conn)

	return &veterinaryService{client}
}

func (s *veterinaryService) SendMakeAppointmentRequest(request *pb.MakeAppointmentRequest) (string, error) {

	ctx := context.Background()

	response, err := s.client.MakeAppointment(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling MakeAppointment")
		return response.Message, err
	}

	return response.Message, err
}

func (s *veterinaryService) SendEndAppointmentRequest(request *pb.EndAppointmentRequest) (string, error) {

	ctx := context.Background()

	response, err := s.client.EndAppointment(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling EndAppointment")
		return response.Message, err
	}

	return response.Message, err
}
