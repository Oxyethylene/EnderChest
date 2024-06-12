package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    int
	Message string
	Data    interface{}
}

func newResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		data, exists := c.Get("response")
		if !exists {
			return
		}
		c.JSON(http.StatusOK, data)
	}
}
