using Microsoft.EntityFrameworkCore;
using System.ComponentModel.DataAnnotations;

namespace OwnerService.Domain.Entities
{
    [Index(nameof(Email), IsUnique = true)]
    public class Owner : Base
    {
        [Required]
        public string Email { get; set; }
        public virtual List<Pet> Pets { get; set; } = new();
    }
}
