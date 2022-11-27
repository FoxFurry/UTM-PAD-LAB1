package config

type Config struct {
	GRPCPort int `env:"GRPC_PORT" envDefault:"24000"`

	DBType string `env:"DB_TYPE" envDefault:"mysql"`
	DBName string `env:"DB_NAME" envDefault:"pad"`
	DBAddr string `env:"DB_ADDR" envDefault:"127.0.0.1:3306"`
	DBUser string `env:"DB_USER" envDefault:"root"`
	DBPass string `env:"DB_PASS" envDefault:"temp"`
}
