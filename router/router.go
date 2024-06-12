package router

import (
	"github.com/Oxyethylene/EnderChest/api"
	"github.com/Oxyethylene/EnderChest/middleware"
	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	g.Use(
		middleware.GinLogger(),
		middleware.GinRecovery(),
		middleware.Cors(),
	)

	userHandler := api.NewLoginApi()
	g.POST("/login", userHandler.Login)

	objectHandler := api.NewObjectApi()
	apiGroup := g.Group("/api")
	apiGroup.Use(middleware.Jwt())
	{
		apiGroup.GET("/object", objectHandler.List)
		apiGroup.PUT("/object", objectHandler.Put)
		apiGroup.GET("/object/:id", objectHandler.Get)
		apiGroup.DELETE("/file/:id", objectHandler.Remove)
	}

	return g
}
