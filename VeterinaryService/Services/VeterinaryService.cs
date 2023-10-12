using Grpc.Core;
using VeterinaryService;
using VeterinaryService.Domain;
using VeterinaryService.Domain.Entities;

namespace VeterinaryService.Services
{
    public class VeterinaryService : Veterinary.VeterinaryBase
    {
        private readonly ILogger<VeterinaryService> _logger;
        private readonly VeterinaryDbContext _veterinaryDbContext;

        public VeterinaryService(ILogger<VeterinaryService> logger, VeterinaryDbContext veterinaryDbContext)
        {
            _logger = logger;
            _veterinaryDbContext = veterinaryDbContext;
        }

        public override Task<MakeAppointmentResponse> MakeAppointment(MakeAppointmentRequest request, ServerCallContext context)
        {
            var appointment = new Appointment
            {
                Id = Guid.NewGuid().ToString(),
                CreatedAt = DateTimeOffset.UtcNow,
                UpdatedAt = DateTimeOffset.UtcNow,
                PetId = request.PetId,
                DateTime = request.DateTime.ToDateTimeOffset(),
            };
            _veterinaryDbContext.Appointments.Add(appointment);
            _veterinaryDbContext.SaveChanges();
            return Task.FromResult(new MakeAppointmentResponse { Message = "Appointment made" });
        }

        public override Task<EndAppointmentResponse> EndAppointment(EndAppointmentRequest request, ServerCallContext context)
        {
            Appointment? appointment = _veterinaryDbContext.Appointments.FirstOrDefault(x => x.Id == request.AppointmentId);

            if (appointment == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Appointment not found"));
            }

            appointment.IsEnded = true;
            appointment.Details = request.Details;
            appointment.UpdatedAt = DateTimeOffset.UtcNow;

            _veterinaryDbContext.Appointments.Update(appointment);
            _veterinaryDbContext.SaveChanges();

            return Task.FromResult(new EndAppointmentResponse { Message = "Appointment ended" });
        }
    }
}