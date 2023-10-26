using Microsoft.EntityFrameworkCore;

namespace OwnerService.Domain
{
    public class OwnerDbContext : DbContext
    {
        public OwnerDbContext() { }
        public OwnerDbContext(DbContextOptions<OwnerDbContext> options) : base(options) { }

        public virtual DbSet<Entities.Owner> Owners {  get; set; }
        public virtual DbSet<Entities.Pet> Pets {  get; set; }
    }
}
