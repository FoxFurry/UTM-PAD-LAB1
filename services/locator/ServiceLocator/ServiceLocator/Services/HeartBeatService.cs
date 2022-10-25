using System;
using System.Collections.Generic;
using System.Net;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Net.Client;
using Microsoft.Extensions.Hosting;
using ServiceLocator.Viewmodels;

namespace ServiceLocator.Services
{
    public class HeartBeatService: BackgroundService
    {
        private RegistryService _registryService;

        public HeartBeatService(RegistryService registryService)
        {
            _registryService = registryService;
        }

        protected override Task ExecuteAsync(CancellationToken stoppingToken)
        {
            Task.Factory.StartNew(() =>
            {
                while (true)
                {
                    HeartBeat();
                    Thread.Sleep(10000);
                }
            });
            return Task.CompletedTask;
        }
        private async void HeartBeat()
        {
            
            foreach (KeyValuePair<ServiceType, List<ServiceModel>> entry in _registryService.registryService)
            {
                List<ServiceModel> toRemove = new List<ServiceModel>();
                foreach (var service in entry.Value)
                {
                    try
                    {
                        Console.WriteLine($"HeartBeat is checking the address {service.Address}");
                        var channel = GrpcChannel.ForAddress(service.Address);
                        var client = new Locator.LocatorClient(channel);
                        
                        if (response.StatusCode == HttpStatusCode.OK)
                        {
                            service.ErrorEpochs = 0;
                        }

                    }
                    catch (Exception e)
                    {
                        Console.WriteLine($"{service.Address} has failed");
                        service.ErrorEpochs++;
                        if (service.ErrorEpochs >= 3)
                        {
                            Console.WriteLine($"Service with address {service.Address} was removed from the list");
                            toRemove.Add(service);
                        }
                    }

                }
                foreach (var serviceModel in toRemove)
                {
                    entry.Value.Remove(serviceModel);
                }
            }
        }
        
        public Task StopAsync(CancellationToken cancellationToken)
        {
            return Task.CompletedTask;
        }
        
    }
}