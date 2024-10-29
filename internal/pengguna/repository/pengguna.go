package repository

import (
	"gorm.io/gorm"
)

type IPenggunaRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type PenggunaRepository struct {
	db *gorm.DB
}

func NewPenggunaRepository(db *gorm.DB) IPenggunaRepository {
	return &PenggunaRepository{
		db: db,
	}
}

