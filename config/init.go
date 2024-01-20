package config

import (
	"container/list"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
)

type configInitializer func()

var initConfigs = list.New()

func InitConfig() {
	viper.SetConfigName("application")
	viper.SetConfigType("toml")
	configPath, _ := filepath.Abs(".")
	viper.AddConfigPath(configPath)
	zap.S().Infow("loading config", "search_path", configPath)
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		zap.S().Fatalw("failed load config",
			"load_path", configPath,
			zap.Error(err),
		)
	}

	for elem := initConfigs.Front(); elem != nil; elem = elem.Next() {
		elem.Value.(configInitializer)()
	}
}
