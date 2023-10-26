using Azure.Core;
using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using OwnerService.Domain;
using OwnerService.Domain.Entities;
using OwnerService.Repos;
using System.Linq;

namespace OwnerService.Services
{
    public class OwnerService : Owner.OwnerBase
    {
        private readonly ILogger<OwnerService> _logger;
        private readonly OwnerDbContext _ownerDbContext;
        private readonly IOwnerRepo _ownerRepo;

        public OwnerService(ILogger<OwnerService> logger, OwnerDbContext ownerDbContext, IOwnerRepo ownerRepo)
        {
            _logger = logger;
            _ownerDbContext = ownerDbContext;
            _ownerRepo = ownerRepo;
        }

        public override async Task<RegisterResponse> Register(RegisterRequest request, ServerCallContext context)
        {
            string id = Guid.NewGuid().ToString();
            var owner = new Domain.Entities.Owner
            {
                Id = id,
                CreatedAt = DateTimeOffset.UtcNow,
                UpdatedAt = DateTimeOffset.UtcNow,
                Email = request.Email,
            };
            _ownerDbContext.Owners.Add(owner);
            await _ownerDbContext.SaveChangesAsync(context.CancellationToken);

            return new RegisterResponse { Message = "Register successful", OwnerId = id };
        }


        public override async Task<RegisterPetResponse> RegisterPet(RegisterPetRequest request, ServerCallContext context)
        {
            var owner = await _ownerRepo.GetOwner(_ownerDbContext, request.OwnerId, request.Email, context.CancellationToken);

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
            await _ownerDbContext.SaveChangesAsync(context.CancellationToken);
            return new RegisterPetResponse { Message = "Pet Register successful" };
        }

        public override async Task<DeleteResponse> Delete(DeleteRequest request, ServerCallContext context)
        {
            var owner = await _ownerRepo.GetOwner(_ownerDbContext, request.OwnerId, request.Email, context.CancellationToken);

            if (owner == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Owner not found"));
            }

            _ownerDbContext.Owners.Remove(owner);
            await _ownerDbContext.SaveChangesAsync(context.CancellationToken);

            return new DeleteResponse { Message = "Owner deleted successfully" };
        }

        public override async Task<GetPetsResponse> GetPets(GetPetsRequest request, ServerCallContext context)
        {
            var owner = await _ownerRepo.GetOwner(_ownerDbContext, request.OwnerId, request.Email, context.CancellationToken);

            if (owner == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Owner not found"));
            }

            var pets = owner.Pets.Select(x => new PetResponse { Name = x.Name, Race = x.Race, Type = x.Type });

            var response = new GetPetsResponse { Message = "ok" };
            response.Pets.AddRange(pets);
            return response;
        }
    }
}