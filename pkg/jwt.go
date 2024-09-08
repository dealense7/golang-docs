package pkg

import (
	"time"

	"github.com/dealense7/market-price-go/config"
	"github.com/dgrijalva/jwt-go"
)

func CreateJwtToken(secret string, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userID,
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "nil", err
	}

	return tokenString, nil
}
