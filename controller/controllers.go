package controller

type Controllers interface {
	Status() StatusChecker
}

type controllers struct {
	status StatusChecker
}

func NewControllers() Controllers {
	return &controllers{
		status: NewStatusController(),
	}
}

func (ctrl *controllers) Status() StatusChecker {
	return ctrl.status
}
