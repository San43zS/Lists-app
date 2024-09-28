package config

import "github.com/spf13/viper"

type Config struct {
	URL    string
	Driver string
}

func NewConfig() Config {
	return Config{
		URL:    viper.GetString("DB_URL"),
		Driver: viper.GetString("DB_DRIVER"),
	}
}
