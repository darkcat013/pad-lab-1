namespace OwnerService.Domain.Entities
{
    public class Pet : Base
    {
        public string Type { get; set; }
        public string Race { get; set; }
        public string Name { get; set; }
        public virtual Owner Owner { get; set; }
    }
}
