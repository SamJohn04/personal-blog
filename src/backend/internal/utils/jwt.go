package utils

import (
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.Cfg.JWTSecret)

func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtKey)
}
