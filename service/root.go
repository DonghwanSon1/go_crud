package service

import (
	"go_crud/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	repository *repository.Repository

	User       *User
	Middleware *Middleware
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}
		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
