package repository

import (
	"gorm.io/gorm"
)

type IRekomendasiNutrisiharianRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type RekomendasiNutrisiHarianRepository struct {
	db *gorm.DB
}

func NewRekomendasiNutrisiHarianRepository(db *gorm.DB) IRekomendasiNutrisiharianRepository {
	return &RekomendasiNutrisiHarianRepository{
		db: db,
	}
}

