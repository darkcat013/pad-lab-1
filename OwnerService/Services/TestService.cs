using Grpc.Core;

namespace OwnerService.Services
{
    public class TestService : Test.TestBase
    {
        private readonly ILogger<TestService> _logger;
        public TestService(ILogger<TestService> logger)
        {
            _logger = logger;
        }

        public override async Task<TestTimeoutResponse> TestTimeout(TestTimeoutRequest request, ServerCallContext context)
        {
            Thread.Sleep(3000);
            return new TestTimeoutResponse();
        }

        public override async Task<TestRateLimitResponse> TestRateLimit(TestRateLimitRequest request, ServerCallContext context)
        {
            return new TestRateLimitResponse();
        }

        public override async Task<TestCircuitBreakerResponse> TestCircuitBreaker(TestCircuitBreakerRequest request, ServerCallContext context)
        {
            throw new Exception();
        }
    }
}
