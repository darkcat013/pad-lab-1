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

        public override Task<TestTimeoutResponse> TestTimeout(TestTimeoutRequest request, ServerCallContext context)
        {
            Thread.Sleep(3000);
            return Task.FromResult(new TestTimeoutResponse());
        }

        public override Task<TestRateLimitResponse> TestRateLimit(TestRateLimitRequest request, ServerCallContext context)
        {
            return Task.FromResult(new TestRateLimitResponse());
        }

        public override Task<TestCircuitBreakerResponse> TestCircuitBreaker(TestCircuitBreakerRequest request, ServerCallContext context)
        {
            return Task.FromResult(new TestCircuitBreakerResponse());
        }
    }
}
