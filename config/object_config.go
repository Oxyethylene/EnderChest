package config

import "github.com/spf13/viper"

var Store StoreConfig

func init() {
	Store = StoreConfig{
		DbPath: viper.GetString("store.path"),
	}
}

type StoreConfig struct {
	DbPath string
}
