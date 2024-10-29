package repository

import (
	"gorm.io/gorm"
)

type IAdministratorRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type AdministratorRepository struct {
	db *gorm.DB
}

func NewAdministratorRepository(db *gorm.DB) IAdministratorRepository {
	return &AdministratorRepository{
		db: db,
	}
}

