package api

import (
	"github.com/Oxyethylene/littlebox/db"
	"github.com/Oxyethylene/littlebox/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type LoginApi struct {
}

func NewLoginApi() *LoginApi {
	return &LoginApi{}
}

type loginParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *LoginApi) Login(c *gin.Context) {
	var param loginParam
	err := c.BindJSON(&param)
	if err != nil {
		zap.S().Fatal("Unmarshall login param failed", "err", err)
	}
	user := db.UserByEmail(param.Email)
	if user.Password != param.Password {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "incorrect username or password",
			"data":    nil,
		})
		return
	}
	token, expireTime, err := util.GenerateToken(user)
	if err != nil {
		zap.S().Error("generate token failed, %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "generate token failed",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "success",
		"data": map[string]interface{}{
			"expired_at": expireTime.Format(time.RFC3339),
			"token":      token,
		},
	})
}
