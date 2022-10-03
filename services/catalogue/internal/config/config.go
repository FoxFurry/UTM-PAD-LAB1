package config

type Config struct {
	GRPCPort int `mapstructure:"grpc_port"`

	DBType string `mapstructure:"db_type"`
	DBName string `mapstructure:"db_name"`
	DBAddr string `mapstructure:"db_addr"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
}
