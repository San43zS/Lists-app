package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	URL    string
	Driver string
}

type dbParams struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func getDBParams() dbParams {
	return dbParams{
		host:     viper.GetString(fmt.Sprintf("%s.host", "db")),
		port:     viper.GetString(fmt.Sprintf("%s.port", "db")),
		user:     viper.GetString(fmt.Sprintf("%s.user", "db")),
		password: viper.GetString(fmt.Sprintf("%s.password", "db")),
		database: viper.GetString(fmt.Sprintf("%s.database", "db")),
	}
}

func (db dbParams) ParseURL() string {
	template := viper.GetString(fmt.Sprintf("%s.urlTemplate", "db"))

	return fmt.Sprintf(template, db.host, db.port, db.database, db.user, db.password)
}

func NewConfig() Config {
	return Config{
		URL:    getDBParams().ParseURL(),
		Driver: viper.GetString(fmt.Sprintf("%s.driver", "db")),
	}
}
