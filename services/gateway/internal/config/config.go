package config

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`

	ServiceAddress string `env:"SERVICE_ADDRESS" envDefault:"localhost:22000"`
	CacheAddress   string `env:"CACHE_ADDRESS" envDefault:"localhost:23000"`
}
