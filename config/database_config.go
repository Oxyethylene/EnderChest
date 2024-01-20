package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var DatabaseConfig = new(databaseConfig)

func init() {
	var f configInitializer = func() {
		if err := viper.UnmarshalKey("datasource", DatabaseConfig); err != nil {
			zap.S().Fatalw("init datasource config failed", zap.Error(err))
		}
	}
	initConfigs.PushBack(f)
}

type databaseConfig struct {
	Driver   string
	Net      string
	Addr     string
	User     string
	Password string
	Database string
}
