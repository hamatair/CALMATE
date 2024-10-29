package repository

import (
	"gorm.io/gorm"
)

type IProgresNutrisiHarianRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
}

type ProgresNutrisiHarianRepository struct {
	db *gorm.DB
}

func NewProgresNutrisiHarianRepository(db *gorm.DB) IProgresNutrisiHarianRepository {
	return &ProgresNutrisiHarianRepository{
		db: db,
	}
}

