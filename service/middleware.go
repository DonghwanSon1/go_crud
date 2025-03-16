package service

import (
	"errors"
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
	token, signErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("유효하지 않은 토큰")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if signErr != nil {
		return "", signErr
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return "", errors.New("유효시간 만료")
		}
		return claims["userId"].(string), nil
	} else {
		return "", err
	}
}
