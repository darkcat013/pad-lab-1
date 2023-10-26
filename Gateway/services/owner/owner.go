package owner

import (
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/darkcat013/pad-lab-1/Gateway/services/owner/pb"
	"github.com/darkcat013/pad-lab-1/Gateway/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OwnerService interface {
	SendRegisterRequest(request *pb.RegisterRequest) (*pb.RegisterResponse, error)
	SendRegisterPetRequest(request *pb.RegisterPetRequest) (string, error)
	SendDeleteRequest(request *pb.DeleteRequest) (string, error)
	SendGetPetsRequest(request *pb.GetPetsRequest) (*pb.GetPetsResponse, error)
}

type ownerService struct {
	client pb.OwnerClient
}

func NewOwnerService(cfg config.Config) OwnerService {
	log.Info().Msg("Creating new owner service")

	conn, err := grpc.Dial(cfg.OwnerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}

	client := pb.NewOwnerClient(conn)

	return &ownerService{client}
}

func (s *ownerService) SendRegisterRequest(request *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	response, err := s.client.Register(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling Register")
		return response, err
	}

	return response, err
}

func (s *ownerService) SendRegisterPetRequest(request *pb.RegisterPetRequest) (string, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	response, err := s.client.RegisterPet(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling RegisterPet")
		return response.Message, err
	}

	return response.Message, err
}

func (s *ownerService) SendDeleteRequest(request *pb.DeleteRequest) (string, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	response, err := s.client.Delete(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling Delete")
		return response.Message, err
	}

	return response.Message, err
}

func (s *ownerService) SendGetPetsRequest(request *pb.GetPetsRequest) (*pb.GetPetsResponse, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	response, err := s.client.GetPets(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling GetPets")
		return response, err
	}

	return response, err
}
