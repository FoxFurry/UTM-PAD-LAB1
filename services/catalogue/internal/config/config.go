package config

type Config struct {
	GRPCPort int `env:"GRPC_PORT" envDefault:"22000"`

	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBName string `env:"DB_NAME" envDefault:"pad"`
	DBAddr string `env:"DB_ADDR" envDefault:"127.0.0.1:3306"`
	DBUser string `env:"DB_USER" envDefault:"root"`
	DBPass string `env:"DB_PASS" envDefault:"temp"`

	MaxConcurrentTasks uint32 `env:"MAX_CONCURRENT_TASKS" envDefault:"16"`

	LocatorAddress string `env:"LOCATOR_ADDRESS" envDefault:"127.0.0.1:5000"`
	CacheAddress   string `env:"CACHE_ADDRESS" envDefault:"127.0.0.1:23000"`
}
