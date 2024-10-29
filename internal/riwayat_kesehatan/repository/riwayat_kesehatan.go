package repository

import (
	"gorm.io/gorm"
)

type IRiwayatKesehatanRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type RiwayatKesehatanRepository struct {
	db *gorm.DB
}

func NewRiwayatKesehatanRepository(db *gorm.DB) IRiwayatKesehatanRepository {
	return &RiwayatKesehatanRepository{
		db: db,
	}
}

