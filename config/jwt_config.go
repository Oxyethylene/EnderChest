package config

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var JwtConfig = new(jwtConfig)

func init() {
	var f configInitializer = func() {
		if err := viper.UnmarshalKey("jwt", JwtConfig); err != nil {
			zap.S().Fatalw("init jwt config failed", zap.Error(err))
		}
	}
	initConfigs.PushBack(f)
}

type jwtConfig struct {
	Secret string
}
