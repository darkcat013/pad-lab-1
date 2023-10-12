using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using OwnerService.Domain;
using OwnerService.Domain.Entities;

namespace OwnerService.Services
{
    public class OwnerService : Owner.OwnerBase
    {
        private readonly ILogger<OwnerService> _logger;
        private readonly OwnerDbContext _ownerDbContext;

        public OwnerService(ILogger<OwnerService> logger, OwnerDbContext ownerDbContext)
        {
            _logger = logger;
            _ownerDbContext = ownerDbContext;
        }

        public override Task<RegisterResponse> Register(RegisterRequest request, ServerCallContext context)
        {
            var owner = new Domain.Entities.Owner
            {
                Id = Guid.NewGuid().ToString(),
                CreatedAt = DateTimeOffset.UtcNow,
                UpdatedAt = DateTimeOffset.UtcNow,
                Email = request.Email,
            };
            _ownerDbContext.Owners.Add(owner);
            _ownerDbContext.SaveChanges();

            return Task.FromResult(new RegisterResponse { Message = "Register successful" });
        }

        public override Task<RegisterPetResponse> RegisterPet(RegisterPetRequest request, ServerCallContext context)
        {
            var owner = _ownerDbContext.Owners.FirstOrDefault(x => x.Id == request.OwnerId);

            if (owner == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Owner not found"));
            }

            owner.Pets.Add(new Pet
            {
                Id = Guid.NewGuid().ToString(),
                CreatedAt = DateTimeOffset.UtcNow,
                UpdatedAt = DateTimeOffset.UtcNow,
                Name = request.Name,
                Race = request.Race,
                Type = request.Type,
            });
            _ownerDbContext.Update(owner);
            _ownerDbContext.SaveChanges();
            return Task.FromResult(new RegisterPetResponse { Message = "Pet Register successful" });
        }

        public override Task<DeleteResponse> Delete(DeleteRequest request, ServerCallContext context)
        {
            var owner = _ownerDbContext.Owners.FirstOrDefault(x => x.Id == request.Id);

            if (owner == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Owner not found"));
            }

            _ownerDbContext.Owners.Remove(owner);
            _ownerDbContext.SaveChanges();

            return Task.FromResult(new DeleteResponse { Message = "Owner deleted successfully" });
        }
    }
}