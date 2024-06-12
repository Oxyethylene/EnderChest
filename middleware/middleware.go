package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return newLog()
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery() gin.HandlerFunc {
	config := RecoveryConfig{
		Stack: true,
	}
	return newRecovery(config)
}

// Cors 处理cors
func Cors() gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowOrigins:     []string{"*"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	return cors.New(config)
}

// Jwt 用JWT鉴权并给接口提供用户信息
func Jwt() gin.HandlerFunc {
	config := JwtConfig{}
	return newJwt(config)
}

func Response() gin.HandlerFunc {
	return newResponse()
}
