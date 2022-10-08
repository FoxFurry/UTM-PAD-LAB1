package config

type Config struct {
	GRPCPort int `env:"GRPC_PORT" envDefault:"22000"`

	ServiceAddress string `env:"SERVICE_ADDRESS" envDefault:"localhost:22000"`
	CacheAddress   string `env:"CACHE_ADDRESS" envDefault:"localhost:23000"`
}
