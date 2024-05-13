package service

import (
	"golang_backend_study/repository"
	"sync"
)

// Network, Repository 다리 역할

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository
	repository *repository.Repository

	User *User
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
