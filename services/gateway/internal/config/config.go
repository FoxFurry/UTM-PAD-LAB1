package config

type Config struct {
	Port int `env:"PORT" envDefault:"8080"`

	CacheAddress   string `env:"CACHE_ADDRESS" envDefault:"127.0.0.1:23000"`
	UserAddress    string `env:"USER_ADDRESS" envDefault:"127.0.0.1:24000"`
	LocatorAddress string `env:"LOCATOR_ADDRESS" envDefault:"127.0.0.1:5000"`
}
