package repository

import (
	"context"

	"go-template/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repository *userRepository) Create(ctx context.Context, user *model.User) error {
	return repository.db.WithContext(ctx).Create(user).Error
}

func (repository *userRepository) Update(ctx context.Context, user *model.User) error {
	return repository.db.WithContext(ctx).Save(user).Error
}

func (repository *userRepository) Delete(ctx context.Context, id uint) error {
	return repository.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

func (repository *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := repository.db.WithContext(ctx).First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *userRepository) List(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if err := repository.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
