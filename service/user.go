package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go_crud/models"
	"go_crud/repository"
	"go_crud/types"
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

func (u *User) Signup(newUser *models.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser.Password = string(hash)
	return u.userRepository.Signup(newUser)
}

func (u *User) Login(req types.LoginRequest) (tokenString string, err error) {

	user, dbErr := u.userRepository.FindUserByEmail(req.Email)
	if dbErr != nil {
		return "", errors.New("해당 사용자가 존재하지 않습니다")
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if bcryptErr != nil {
		return "", errors.New("비밀번호가 일치하지 않습니다")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(30 * time.Minute).Unix(), // 30분
	})

	tokenString, err = token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", errors.New("JWT Signing 에러입니다")
	}

	return tokenString, nil
}

func (u *User) Update(name string, newAge int64) error {
	return u.userRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user.Email)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
