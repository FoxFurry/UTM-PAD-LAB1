package config

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`

	CacheAddress   string `env:"CACHE_ADDRESS" envDefault:"localhost:23000"`
	UserAddress    string `env:"USER_ADDRESS" envDefault:"localhost:24000"`
	LocatorAddress string `env:"LOCATOR_ADDRESS" envDefault:"localhost:5000"`
}
