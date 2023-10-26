using Microsoft.EntityFrameworkCore;
using OwnerService.Domain;
using OwnerService.Repos;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc();
// Initialize EF Core
builder.Services.AddDbContext<OwnerDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DbConnection"))
            .UseLazyLoadingProxies());
builder.Services.AddScoped<IOwnerRepo, OwnerRepo>();

var app = builder.Build();

//Ensure DB is created and seed it
using (var scope = app.Services.CreateScope())
{
    var services = scope.ServiceProvider;

    var context = services.GetRequiredService<OwnerDbContext>();

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
app.MapGrpcService<OwnerService.Services.OwnerService>();
app.MapGrpcService<OwnerService.Services.NotificationService>();
app.MapGrpcService<OwnerService.Services.TestService>();

app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client.");

app.MapGet("/status", () => new { status = "ok" });

app.Run();
