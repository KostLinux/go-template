package service

import "go-template/repository"

type Services interface {
	User() UserService
}

type services struct {
	user UserService
}

func NewServices(repos repository.Repositories) Services {
	return &services{
		user: NewUserService(repos.User()),
	}
}

func (service *services) User() UserService {
	return service.user
}
