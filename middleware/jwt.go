package middleware

import (
	"fmt"
	"github.com/Oxyethylene/littlebox/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// JwtConfig defines the config for Jwt middleware
type JwtConfig struct {
	TokenLookUp string
	// TokenHeadName is a string prefix in the header. Default value is "Bearer "
	TokenHeaderName string
}

// newJwt is a middleware that checks for a JWT on the `Authorization` header
func newJwt(config JwtConfig) gin.HandlerFunc {
	if config.TokenHeaderName == "" {
		config.TokenHeaderName = "Bearer "
	}
	if config.TokenLookUp == "" {
		config.TokenLookUp = "header:Authorization"
	}
	return func(c *gin.Context) {
		code := 401

		source, key, _ := strings.Cut(config.TokenLookUp, ":")
		token := ""
		if source == "header" {
			token = lookupTokenFromHeader(key, c)
		}
		token = strings.TrimPrefix(token, config.TokenHeaderName)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  fmt.Sprintf("No Auth info in %s '%s'", source, key),
			})

			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "Invalid Token",
			})

			c.Abort()
			return
		} else if time.Now().After(claims.ExpiresAt.Time) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "Token Expired",
			})

			c.Abort()
			return
		}

		c.Set("User", claims)

		c.Next()
	}
}

func lookupTokenFromHeader(header string, c *gin.Context) string {
	return c.GetHeader(header)
}
