namespace ServiceLocator.Viewmodels
{
    public class ServiceModel
    {
        public string Address { get; set; }
        public int ErrorEpochs { get; set; }

        public ServiceModel(string address)
        {
            Address = address;
            ErrorEpochs = 0;
        }
    }
}