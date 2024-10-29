package repository

import (
	"gorm.io/gorm"
)

type IArtikelRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type ArtikelRepository struct {
	db *gorm.DB
}

func NewArtikelRepository(db *gorm.DB) IArtikelRepository {
	return &ArtikelRepository{
		db: db,
	}
}

