package controller

import "go-template/service"

type Controllers interface {
	User() User
}

type controllers struct {
	user User
}

func NewControllers(services service.Services) Controllers {
	return &controllers{
		user: NewUserController(services.User()),
	}
}

func (ctrl *controllers) User() User {
	return ctrl.user
}
