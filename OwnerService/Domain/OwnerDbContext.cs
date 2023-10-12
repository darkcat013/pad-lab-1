using Microsoft.EntityFrameworkCore;

namespace OwnerService.Domain
{
    public class OwnerDbContext : DbContext
    {
        public OwnerDbContext(DbContextOptions<OwnerDbContext> options) : base(options) { }

        public DbSet<Entities.Owner> Owners {  get; set; }
        public DbSet<Entities.Pet> Pets {  get; set; }
    }
}
