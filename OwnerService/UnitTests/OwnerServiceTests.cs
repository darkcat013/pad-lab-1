using Grpc.Core;
using Microsoft.EntityFrameworkCore;
using Moq;
using NUnit.Framework;
using OwnerService;
using OwnerService.Domain;
using OwnerService.Repos;

namespace UnitTests
{
    [TestFixture]
    public class OwnerServiceTests
    {
        private OwnerDbContext _ownerDbContext;
        private OwnerService.Services.OwnerService _ownerService;
        private Mock<DbSet<OwnerService.Domain.Entities.Owner>> _ownerDbSet;
        private Mock<OwnerDbContext> _dbContextMock;
        private Mock<ILogger<OwnerService.Services.OwnerService>> _loggerMock;
        private Mock<IOwnerRepo> _repoMock;

        [SetUp]
        public void Setup()
        {
            _ownerDbSet = new Mock<DbSet<OwnerService.Domain.Entities.Owner>>();
            _dbContextMock = new Mock<OwnerDbContext>();
            _dbContextMock.Setup(x => x.Owners).Returns(_ownerDbSet.Object);
            _ownerDbContext = _dbContextMock.Object;
            _loggerMock = new Mock<ILogger<OwnerService.Services.OwnerService>>();
            _repoMock = new Mock<IOwnerRepo>();

            _ownerService = new OwnerService.Services.OwnerService(_loggerMock.Object, _ownerDbContext, _repoMock.Object);
        }

        [Test]
        public async Task Register_OwnerRegistration_Success()
        {
            // Arrange
            var request = new RegisterRequest { Email = "test@example.com" };
            var owner = new OwnerService.Domain.Entities.Owner {Email = request.Email };
            var cancellationToken = new CancellationToken();

            _dbContextMock.Setup(x => x.Owners.Add(It.IsAny<OwnerService.Domain.Entities.Owner>()));
            _dbContextMock.Setup(x => x.SaveChangesAsync(cancellationToken)).ReturnsAsync(1);

            // Act
            var response = await _ownerService.Register(request, GetServerCallContext());

            // Assert
            Assert.That(response.Message, Is.EqualTo("Register successful"));
            Assert.That(Guid.TryParse(response.OwnerId, out var res), Is.EqualTo(true));
        }

        [Test]
        public async Task RegisterPet_PetRegistration_Success()
        {
            // Arrange
            var request = new RegisterPetRequest { OwnerId = "1", Name = "Fido" };
            var owner = new OwnerService.Domain.Entities.Owner { Id = request.OwnerId, Email = "test@example.com" };
            var cancellationToken = new CancellationToken();

            _repoMock.Setup(x => x.GetOwner(It.IsAny<OwnerDbContext>(), It.IsAny<string>(), It.IsAny<string>(), It.IsAny<CancellationToken>())).ReturnsAsync(owner);
            _dbContextMock.Setup(x => x.Update(It.IsAny<OwnerService.Domain.Entities.Owner>()));
            _dbContextMock.Setup(x => x.SaveChangesAsync(cancellationToken)).ReturnsAsync(1);

            // Act
            var response = await _ownerService.RegisterPet(request, GetServerCallContext());

            // Assert
            Assert.That(response.Message, Is.EqualTo("Pet Register successful"));
        }

        // Add similar tests for other methods like Delete and GetPets

        private ServerCallContext GetServerCallContext()
        {
            // Create a mock ServerCallContext for testing
            var callContext = new Mock<ServerCallContext>();
            return callContext.Object;
        }
    }
}