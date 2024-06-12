package main

import (
	"github.com/Oxyethylene/EnderChest/config"
	"github.com/Oxyethylene/EnderChest/db"
	"github.com/Oxyethylene/EnderChest/logging"
	"github.com/Oxyethylene/EnderChest/router"
	"github.com/Oxyethylene/EnderChest/runner"
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
