package config

import "github.com/spf13/viper"

var Database DatabaseConfig

func init() {
	Database = DatabaseConfig{
		Driver:   viper.GetString("datasource.driver"),
		Net:      viper.GetString("datasource.net"),
		Addr:     viper.GetString("datasource.address"),
		User:     viper.GetString("datasource.user"),
		Password: viper.GetString("datasource.password"),
		Database: viper.GetString("datasource.database"),
	}
}

type DatabaseConfig struct {
	Driver   string
	Net      string
	Addr     string
	User     string
	Password string
	Database string
}
