namespace VeterinaryService.Domain.Entities
{
    public class Appointment : Base
    {
        public DateTimeOffset DateTime { get; set; }
        public bool IsEnded { get; set; }
        public string PetId { get; set; }
        public string? Details {  get; set; }
    }
}
