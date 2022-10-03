package config

type Config struct {
	DBName string `mapstructure:"db_name"`
	DBHost string `mapstructure:"db_host"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBPort string `mapstructure:"db_port"`
}
