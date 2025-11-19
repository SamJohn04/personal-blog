package utils

import (
	"time"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(config.Cfg.JWTSecret)

// Generate a signed JWT string with an expiry set to 24 hours from now.
func GenerateJWT(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtKey)
}
