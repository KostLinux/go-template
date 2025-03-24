package service

import (
	"context"
	"errors"

	"go-template/model"
	constants "go-template/pkg/const"
	"go-template/repository"

	"gorm.io/gorm"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (service *userService) Create(ctx context.Context, user *model.User) error {
	// Add business logic here (e.g., validation, password hashing)
	return service.repo.Create(ctx, user)
}

func (service *userService) Update(ctx context.Context, user *model.User) error {
	if _, err := service.GetByID(ctx, user.ID); err != nil {
		return err
	}

	return service.repo.Update(ctx, user)
}

func (service *userService) Delete(ctx context.Context, id uint) error {
	if _, err := service.GetByID(ctx, id); err != nil {
		return err
	}

	return service.repo.Delete(ctx, id)
}

func (service *userService) GetByID(ctx context.Context, id uint) (*model.User, error) {
	// Add business logic here (e.g., caching)
	user, err := service.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, constants.ErrNotFound
		}
		return nil, err
	}

	return user, nil
}

func (service *userService) List(ctx context.Context) ([]model.User, error) {
	return service.repo.List(ctx)
}
