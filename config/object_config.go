package config

import "github.com/spf13/viper"

var Store storeConfig

func init() {
	var f configInitializer = func() {
		Store = storeConfig{
			DbPath: viper.GetString("store.path"),
		}
	}
	initConfigs.PushBack(f)
}

type storeConfig struct {
	DbPath string
}
