package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ServerConfig = new(serverConfig)

func init() {
	var f configInitializer = func() {
		viper.SetDefault("application.address", "0.0.0.0")
		viper.SetDefault("application.port", 8080)
		err := viper.UnmarshalKey("application", ServerConfig)
		if err != nil {
			zap.S().Fatalw("parse application config failed", zap.Error(err))
		}
	}
	initConfigs.PushBack(f)
}

type serverConfig struct {
	Addr string `mapstructure:"address"`
	Port int    `mapstructure:"port"`
}
