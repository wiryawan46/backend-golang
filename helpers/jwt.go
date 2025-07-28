package helpers

import (
	"backend-golang/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte(config.GetEnv("JWT_KEY", "secret"))

func GenerateToken(username string) string {
	expiresAt := time.Now().Add(60 * time.Minute)
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}
