package service

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go_crud/models"
	"go_crud/repository"
	"go_crud/types"
	"go_crud/types/usersInfo"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Signup(newUser *models.UsersInfo) error {
	// 비밀번호 인코딩
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)
	return u.userRepository.Signup(newUser)
}

func (u *User) Login(req usersInfo.LoginRequest) (tokenString string, err error) {

	user, dbErr := u.userRepository.FindUserByEmail(req.UserId)
	if dbErr != nil {
		return "", errors.New("해당 사용자가 존재하지 않습니다")
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if bcryptErr != nil {
		return "", errors.New("비밀번호가 일치하지 않습니다")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserId,
		"exp":    time.Now().Add(30 * time.Minute).Unix(), // 30분
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("JWT Signing 에러입니다")
	}

	return tokenString, nil
}

func (u *User) RefreshToken(tokenString string) (refreshTokenString string, err error) {
	beforeToken, signErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("유효하지 않은 토큰")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if signErr != nil {
		return "", signErr
	}

	if claims, ok := beforeToken.Claims.(jwt.MapClaims); ok {
		// 토큰 시간이 만료 시 토큰 재발급
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			user, dbErr := u.userRepository.FindUserByEmail(claims["userId"].(string))
			if dbErr != nil {
				return "", errors.New("해당 사용자가 존재하지 않습니다")
			}

			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"userId": user.UserId,
				"exp":    time.Now().Add(30 * time.Minute).Unix(), // 30분
			})

			refreshTokenString, err = refreshToken.SignedString([]byte(os.Getenv("SECRET")))
			if err != nil {
				return "", errors.New("JWT Signing 에러입니다")
			}

			return refreshTokenString, nil
		}
		// 만료시간이 끝나지 않았을 시 그대로 토큰 리턴
		return tokenString, nil
	} else {
		return "", err
	}
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Email)
}

func (u *User) Get(userId string) []*types.User {
	fmt.Println(userId)
	return u.userRepository.Get()
}
