package utils

import (
	"fmt"
	"go-fiber-app/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(models *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username":   models.Username,
		"FirstName":  models.FirstName,
		"LastName":   models.LastName,
		"Email":      models.Email,
		"ExpireTime": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func DecodeToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

