package database

type DBConfig struct {
	Type string `mapstructure:"db_type"`
	Name string `mapstructure:"db_name"`
	Host string `mapstructure:"db_host"`
	User string `mapstructure:"db_user"`
	Pass string `mapstructure:"db_pass"`
	Port string `mapstructure:"db_port"`
}
