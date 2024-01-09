package util

import (
	"github.com/Oxyethylene/littlebox/db"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("+fHCDHCRBbd=B9,~-_s=VyF:%?P:}E5HGWtwBX4@J#cDgxaMBeifQ@_6?spP5@*R")

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *db.User) (string, time.Time, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)

	claims := Claims{
		user.Id,
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(nowTime),
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, expireTime, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
