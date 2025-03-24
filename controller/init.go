package controller

import "go-template/service"

type Controllers interface {
	Status() StatusChecker
	User() User
}

type controllers struct {
	status StatusChecker
	user   User
}

func NewControllers(services service.Services) Controllers {
	return &controllers{
		status: NewStatusController(),
		user:   NewUserController(services.User()),
	}
}

func (ctrl *controllers) Status() StatusChecker {
	return ctrl.status
}

func (ctrl *controllers) User() User {
	return ctrl.user
}
