using System.Threading.Tasks;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;
using Microsoft.Extensions.Logging;

namespace ServiceLocator.Services
{
    public class LocatorService: Locator.LocatorBase
    {
        private RegistryService _registryService;
        private readonly ILogger<LocatorService> _logger;

        public LocatorService(ILogger<LocatorService> logger,
            RegistryService registryService)
        {
            _registryService = registryService;
            _logger = logger;
        }

        public override Task<GetServiceResponse> GetService(GetServiceRequest request, ServerCallContext context)
        {
            var response = _registryService.GetService(request);
            if (response == null)
            {
                throw new RpcException(new Status(StatusCode.NotFound, "Not Found"), new Metadata()
                {
                    {"ServiceType", request.Type.ToString()}
                });
            }

            GetServiceResponse result = new GetServiceResponse()
            {
                Address = response.Address
            };
            return Task.FromResult(result);
        }

        public override Task<Empty> RegisterService(RegisterServiceRequest request, ServerCallContext context)
        {
            _registryService.RegisterService(request);
            return Task.FromResult(new Empty());
        }
    }
}