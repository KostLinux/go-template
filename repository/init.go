package repository

import "gorm.io/gorm"

type Repositories interface {
	User() UserRepository
}

type repositories struct {
	db   *gorm.DB
	user UserRepository
}

func NewRepositories(db *gorm.DB) Repositories {
	return &repositories{
		db:   db,
		user: NewUserRepository(db),
	}
}

func (r *repositories) User() UserRepository {
	return r.user
}
