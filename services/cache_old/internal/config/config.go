package config

type Config struct {
	GRPCPort int `env:"GRPC_PORT" envDefault:"23000"`
}
