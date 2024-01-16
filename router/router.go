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
		middleware.Cors(),
	)

	userHandler := api.NewLoginApi()
	g.POST("/login", userHandler.Login)

	objectHandler := api.NewObjectApi()
	apiGroup := g.Group("/api")
	apiGroup.Use(middleware.Jwt())
	{
		apiGroup.GET("/file", objectHandler.List)
		apiGroup.POST("file", objectHandler.Add)
		apiGroup.GET("/file/:id", objectHandler.Get)
		apiGroup.DELETE("/file/:id", objectHandler.Remove)
	}

	return g
}
