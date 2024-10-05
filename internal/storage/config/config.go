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
	dbname   string
}

func getDBParams() *dbParams {
	return &dbParams{
		host:     viper.GetString("HOST"),
		port:     viper.GetString("PORT"),
		user:     viper.GetString("USER"),
		password: viper.GetString("PASSWORD"),
		dbname:   viper.GetString("DBNAME"),
	}
}

func (db dbParams) ParseURL() string {
	template := viper.GetString("URLTEMPLATE")

	return fmt.Sprintf(template, db.host, db.port, db.dbname, db.user, db.password)
}

func NewConfig() *Config {
	test := getDBParams()

	return &Config{
		URL:    test.ParseURL(),
		Driver: viper.GetString("DRIVER"),
	}
}
