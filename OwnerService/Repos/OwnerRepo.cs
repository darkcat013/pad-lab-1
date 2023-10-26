using Microsoft.EntityFrameworkCore;
using OwnerService.Domain;

namespace OwnerService.Repos
{
    public class OwnerRepo : IOwnerRepo
    {
        public async Task<Domain.Entities.Owner> GetOwner(OwnerDbContext ownerDbContext, string id, string email, CancellationToken cancellationToken)
        {
            return await ownerDbContext.Owners.FirstOrDefaultAsync(x => x.Id == id || x.Email.Equals(email), cancellationToken);
        }
    }
}
