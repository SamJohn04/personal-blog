package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/SamJohn04/personal-blog/src/backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

const (
	userIdKey   = "userId"
	userAuthKey = "userAuth"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id, authLevel int

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWTSecret), nil
		})
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusBadGateway)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			http.Error(w, "Something went wrong", http.StatusBadGateway)
			return
		}

		err = config.DB.QueryRow("SELECT id, auth_level FROM users WHERE email=?", email).Scan(&id, &authLevel)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(
			context.WithValue(r.Context(), userIdKey, id),
			userAuthKey,
			authLevel,
		)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserAuth(r *http.Request) (int, error) {
	authLevel, ok := r.Context().Value(userAuthKey).(int)
	if !ok {
		return 0, errors.New("not a valid auth level")
	}
	return authLevel, nil
}
