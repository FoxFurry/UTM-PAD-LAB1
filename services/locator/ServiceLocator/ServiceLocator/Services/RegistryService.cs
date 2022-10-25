using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.Threading;
using Grpc.Core;
using ServiceLocator.Viewmodels;

namespace ServiceLocator.Services
{
    public class RegistryService
    {
        private int index = 0;
        public ConcurrentDictionary<ServiceType, List<ServiceModel>> registryService =
            new ConcurrentDictionary<ServiceType, List<ServiceModel>>();

        public RegistryService()
        {
            registryService.TryAdd(ServiceType.Cache, new List<ServiceModel>());
            registryService.TryAdd(ServiceType.Catalogue, new List<ServiceModel>());
        }
        

        public void RegisterService(RegisterServiceRequest request)
        {
            try
            {
                ServiceModel temp = new ServiceModel(request.Address);
                registryService[request.Type].Add(temp);
                Console.WriteLine($"Service with address {request.Address} was successfully added");
            }
            catch(Exception e)
            {
                Console.WriteLine("Could not register service");
                throw new RpcException(new Status(StatusCode.Internal, "Internal"), new Metadata()
                { });
            }
        }

        public ServiceModel GetService(GetServiceRequest request)
        {
            if (registryService[request.Type].Count == 0)
            {
                Console.WriteLine($"No services found for {request.Type}");
                throw new RpcException(new Status(StatusCode.NotFound, "Not Found"), new Metadata()
                { }); 
            }

            try
            {
                ServiceModel service = registryService[request.Type][index % registryService[request.Type].Count];
                Interlocked.Increment(ref index);
                Console.WriteLine($"Service with address {service.Address} was returned");
                return service;
            }
            catch (Exception e)
            {
                Console.WriteLine("Could not return service");
                throw new RpcException(new Status(StatusCode.Internal, "Internal"), new Metadata()
                { });
            }
        }
    }
}