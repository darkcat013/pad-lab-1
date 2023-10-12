namespace OwnerService.Domain.Entities
{
    public abstract class Base
    {
        public string Id { get; set; }
        public DateTimeOffset CreatedAt { get; set; }
        public DateTimeOffset UpdatedAt { get; set; }
    }
}
