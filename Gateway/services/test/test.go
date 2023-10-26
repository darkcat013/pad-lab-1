package test

import (
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/darkcat013/pad-lab-1/Gateway/services/test/pb"
	"github.com/darkcat013/pad-lab-1/Gateway/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TestService interface {
	TestTimeout(request *pb.TestTimeoutRequest) (string, error)
	TestRateLimit(request *pb.TestRateLimitRequest) (string, error)
	TestCircuitBreaker(request *pb.TestCircuitBreakerRequest) (string, error)
}

type testService struct {
	client pb.TestClient
}

func NewTestService(cfg config.Config) TestService {
	log.Info().Msg("Creating new test service")

	conn, err := grpc.Dial(cfg.OwnerUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}

	client := pb.NewTestClient(conn)

	return &testService{client}
}

func (s *testService) TestTimeout(request *pb.TestTimeoutRequest) (string, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	_, err := s.client.TestTimeout(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling TestTimeout")
		return "Timeout achieved", err
	}

	return "Timeout not working", err
}

func (s *testService) TestRateLimit(request *pb.TestRateLimitRequest) (string, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	_, err := s.client.TestRateLimit(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling TestRateLimit")
		return "Rate limit achieved", err
	}

	return "Rate limit test", err
}

func (s *testService) TestCircuitBreaker(request *pb.TestCircuitBreakerRequest) (string, error) {

	ctx, cancel := utils.GetDeadlineContext()
	defer cancel()

	_, err := s.client.TestCircuitBreaker(ctx, request)
	if err != nil {
		log.Error().Err(err).Msg("Error calling TestCircuitBreaker")
		return "Server Error", err
	}

	return "ok", err
}
