using Microsoft.EntityFrameworkCore;
using OwnerService.Domain;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc();
// Initialize EF Core
builder.Services.AddDbContext<OwnerDbContext>(options =>
    options.UseSqlServer(builder.Configuration.GetConnectionString("DbConnection"))
            .UseLazyLoadingProxies());

var app = builder.Build();

//Ensure DB is created and seed it
using (var scope = app.Services.CreateScope())
{
    var services = scope.ServiceProvider;

    var context = services.GetRequiredService<OwnerDbContext>();
    context.Database.EnsureCreated();
    // DbInitializer.Initialize(context);
}

//Add gRPC
app.MapGrpcService<OwnerService.Services.OwnerService>();
app.MapGet("/", () => "Communication with gRPC endpoints must be made through a gRPC client.");

app.MapGet("/status", () => new { status = "ok" });

app.Run();
