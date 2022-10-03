package config

type Config struct {
	Type string `mapstructure:"db_type"`
	Name string `mapstructure:"db_name"`
	Addr string `mapstructure:"db_addr"`
	User string `mapstructure:"db_user"`
	Pass string `mapstructure:"db_pass"`
}
