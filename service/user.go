package service

import (
	"go_crud/repository"
	"go_crud/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(beforeUser *types.User, updateUser *types.User) error {
	return u.userRepository.Update(beforeUser, updateUser)
}

func (u *User) Delete(user *types.User) error {
	return u.userRepository.Delete(user)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
