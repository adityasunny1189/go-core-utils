package jwt

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateNewToken(key string, value any, expiryTime int, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			key:   value,
			"exp": time.Now().Add(time.Hour * time.Duration(expiryTime)).Unix(),
		})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println("Error: ", err)
		return "", err
	}
	log.Println("Token: ", tokenString)
	return "Bearer " + tokenString, nil
}

func VerifyToken(key string, tokenString string, secretKey string) (any, error) {
	actualToken := strings.Split(tokenString, " ")
	token, err := jwt.Parse(actualToken[1], func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("INVALID TOKEN")
	}

	value := token.Claims.(jwt.MapClaims)[key]
	return value, nil
}
