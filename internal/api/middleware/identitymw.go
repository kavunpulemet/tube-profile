package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
	"tube-profile/internal/utils"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "UserID"
)

var jwtSecret = []byte("HUdsufs&7fgd9Udkf0fsif8f89wFD9Dfef8D9wE#ie")

func UserIdentity(ctx utils.MyContext, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.URL.Path == "/swagger/" || len(r.URL.Path) >= 9 && r.URL.Path[:9] == "/swagger/" {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get(authorizationHeader)
		if authHeader == "" {
			utils.NewErrorResponse(ctx, w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			utils.NewErrorResponse(ctx, w, "invalid auth header format", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]

		claims, err := ParseToken(tokenString)
		if err != nil {
			utils.NewErrorResponse(ctx, w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "invalid token payload: missing user_id", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtx, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("missing 'exp' field in token")
	}

	if time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token has expired")
	}

	return claims, nil
}
