using OwnerService.Domain;

namespace OwnerService.Repos
{
    public interface IOwnerRepo
    {
        public Task<Domain.Entities.Owner> GetOwner(OwnerDbContext ownerDbContext, string id, string email, CancellationToken cancellationToken);
    }
}
