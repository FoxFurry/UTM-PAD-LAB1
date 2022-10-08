package config

type Config struct {
	GRPCPort int `env:"GRPC_PORT" envDefault:"22000"`

	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBName string `env:"DB_NAME" envDefault:"catalogue"`
	DBAddr string `env:"DB_ADDR" envDefault:"localhost:3306"`
	DBUser string `env:"DB_USER" envDefault:"root"`
	DBPass string `env:"DB_PASS" envDefault:"temp"`

	CacheAddress string `env:"CACHE_ADDRESS" envDefault:"localhost:23000"`
}
