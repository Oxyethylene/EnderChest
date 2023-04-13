package main

import (
	"github.com/Oxyethylene/littlebox/logging"
	"github.com/Oxyethylene/littlebox/router"
	"github.com/Oxyethylene/littlebox/runner"
	"go.uber.org/zap"
)

func init() {
	logging.InitLogger()
}

func main() {
	defer zap.S().Sync()

	engine := router.Create()

	runner.Run(engine)
}
