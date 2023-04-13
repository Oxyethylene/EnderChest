package router

import (
	"github.com/Oxyethylene/littlebox/api"
	"github.com/Oxyethylene/littlebox/middleware"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()
	g.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(true),
		middleware.GinCors(),
	)

	objectHandler := api.NewObjectApi()

	g.GET("/files", objectHandler.List)
	g.POST("file", objectHandler.Add)
	g.GET("/file", objectHandler.Get)

	return g
}
