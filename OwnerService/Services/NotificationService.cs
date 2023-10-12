using Grpc.Core;

namespace OwnerService.Services
{
    public class NotificationService : Notification.NotificationBase
    {
        private readonly ILogger<NotificationService> _logger;
        public NotificationService(ILogger<NotificationService> logger)
        {
            _logger = logger;
        }

        public override Task<NotificationResponse> SendNotification(NotificationRequest request, ServerCallContext context)
        {
            return Task.FromResult(new NotificationResponse { Message = "Notifications Sent" });
        }
    }
}
