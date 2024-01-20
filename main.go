package main

import (
	"github.com/Oxyethylene/littlebox/config"
	"github.com/Oxyethylene/littlebox/db"
	"github.com/Oxyethylene/littlebox/logging"
	"github.com/Oxyethylene/littlebox/router"
	"github.com/Oxyethylene/littlebox/runner"
	"go.uber.org/zap"
)

func init() {
	logging.InitLogger()
	config.InitConfig()
	db.InitDbClient()
}

func main() {
	defer func() {
		_ = zap.S().Sync()
	}()

	engine := router.Create()

	runner.Run(engine)
}
