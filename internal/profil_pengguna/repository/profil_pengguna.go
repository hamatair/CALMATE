package repository

import (
	"gorm.io/gorm"
)

type IProfilPenggunaRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type ProfilPenggunaRepository struct {
	db *gorm.DB
}

func NewProfilPenggunaRepository(db *gorm.DB) IProfilPenggunaRepository {
	return &ProfilPenggunaRepository{
		db: db,
	}
}

