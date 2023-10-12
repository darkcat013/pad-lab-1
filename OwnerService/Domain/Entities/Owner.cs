namespace OwnerService.Domain.Entities
{
    public class Owner : Base
    {
        public string Email { get; set; }
        public virtual List<Pet> Pets { get; set; } = new();
    }
}
