package repository

import (
	"go_crud/config"
	"go_crud/models"
	"go_crud/types"
	"go_crud/types/errors"
)

type UserRepository struct {
	userMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		userMap: []*types.User{},
	}
}

func (u *UserRepository) Signup(newUser *models.User) error {
	result := config.DB.Create(newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserRepository) Update(email string, newAge int64) error {
	isExisted := false
	for _, user := range u.userMap {
		if user.Email == email {
			user.Age = newAge
			isExisted = true
			break
		}
	}

	if !isExisted {
		return errors.ErrOf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Delete(userEmail string) error {
	isExisted := false
	for index, user := range u.userMap {
		if user.Email == userEmail {
			u.userMap = append(u.userMap[:index], u.userMap[index+1:]...)
			isExisted = true
			break
		}
	}

	if !isExisted {
		return errors.ErrOf(errors.NotFoundUser, nil)
	} else {
		return nil
	}
}

func (u *UserRepository) Get() []*types.User {
	return u.userMap
}
