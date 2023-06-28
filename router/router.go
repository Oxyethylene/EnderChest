package router

import (
	"github.com/Oxyethylene/littlebox/api"
	"github.com/Oxyethylene/littlebox/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.Cors(),
	)

	authorized := g.Group("/", gin.BasicAuth(gin.Accounts{
		"admin":   "159632",
		"kudlife": "kudlife",
	}))

	objectHandler, err := api.NewObjectApi()
	if err != nil {
		zap.S().Fatalw("error init objectHandler",
			zap.Error(err),
		)
	}

	g.GET("/files", objectHandler.List)
	authorized.POST("/file", objectHandler.Add)
	authorized.GET("/file", objectHandler.Get)
	authorized.DELETE("/file", objectHandler.Remove)

	return g
}
