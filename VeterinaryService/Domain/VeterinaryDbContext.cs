using Microsoft.EntityFrameworkCore;

namespace VeterinaryService.Domain
{
    public class VeterinaryDbContext : DbContext
    {
        public VeterinaryDbContext(DbContextOptions<VeterinaryDbContext> options) : base(options) { }

        public DbSet<Entities.Appointment> Appointments { get; set; }
    }
}
