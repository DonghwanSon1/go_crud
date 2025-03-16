package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type Middleware struct {
}

func newMiddlewareService() *Middleware {
	return &Middleware{}
}

func (m *Middleware) ValidateToken(tokenString string) (userId string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("유효시간 만료")
		}
		return claims["userId"].(string), nil
	} else {
		return "", err
	}
}
