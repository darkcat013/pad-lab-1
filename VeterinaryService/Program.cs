using VeterinaryService.Domain;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc();

// Initialize EF Core
builder.Services.AddDbContext<VeterinaryDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DbConnection"))
    .UseLazyLoadingProxies());


var app = builder.Build();

//Ensure DB is created and seed it
using (var scope = app.Services.CreateScope())
{
    var services = scope.ServiceProvider;

    var context = services.GetRequiredService<VeterinaryDbContext>();
    try
    {
        context.Database.EnsureCreated();
    }
    catch (Exception ex)
    {
        Console.WriteLine($"{ex.Message}");
    }
}

//Add gRPC
app.MapGrpcService<VeterinaryService.Services.VeterinaryService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client.");

app.MapGet("/status", () => new { status = "ok" });

app.Run();
