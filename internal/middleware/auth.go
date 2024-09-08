package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dealense7/market-price-go/config"
	"github.com/dealense7/market-price-go/internal/interfaces"
	"github.com/dealense7/market-price-go/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderRequired      = errors.New("authorization header is required")
	ErrInvalidAuthHeaderFormat = errors.New("invalid authorization header format")
	ErrInvalidToken            = errors.New("invalid token")
	ErrExpiredToken            = errors.New("token has expired")
	ErrUserNotFound            = errors.New("unauthorized")
	ErrInvalidTokenClaims      = errors.New("invalid token claims")
)

func JWTAuthMiddleware(userService interfaces.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the Authorization header (e.g., "Bearer <token>")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.WriteError(c, http.StatusUnauthorized, ErrAuthHeaderRequired)
			c.Abort()
			return
		}

		// Extract the token part from the header (strip the "Bearer" part)
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			utils.WriteError(c, http.StatusUnauthorized, ErrInvalidAuthHeaderFormat)
			c.Abort()
			return
		}

		// Parse the JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.Envs.JWTSecret), nil
		})

		if err != nil {
			utils.WriteError(c, http.StatusUnauthorized, ErrInvalidToken)
			c.Abort()
			return
		}

		// Validate the token claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expiresAt := int64(claims["expiresAt"].(float64))
			if time.Now().Unix() > expiresAt {
				utils.WriteError(c, http.StatusUnauthorized, ErrExpiredToken)
				c.Abort()
				return
			}

			// Set the user ID in the request context
			userID := int(claims["userId"].(float64))
			item, err := userService.GetById(userID)

			if err != nil {
				utils.WriteError(c, http.StatusUnauthorized, ErrUserNotFound)
				c.Abort()
				return
			}

			c.Set("user", item)

			// Proceed to the next handler
			c.Next()
		} else {
			utils.WriteError(c, http.StatusUnauthorized, ErrInvalidTokenClaims)
			c.Abort()
		}
	}
}
