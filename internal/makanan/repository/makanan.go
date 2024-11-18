package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"gorm.io/gorm"
)

type IMakananRepository interface {
	CreateMakanan(entity.Makanan) error
	GetMakanan(string) ([]entity.Makanan, error)
}

type MakananRepository struct {
	db *gorm.DB
}

// GetMakanan implements IMakananRepository.
func (r *MakananRepository) GetMakanan(nama string) ([]entity.Makanan, error) {
	var makanan []entity.Makanan
	err := r.db.Debug().Where("nama = ?", nama).Find(&makanan).Error
	if err != nil {
		return []entity.Makanan{}, err
	}

	return makanan, err
}

// CreateMakanan implements IMakananRepository.
func (r *MakananRepository) CreateMakanan(makanan entity.Makanan) error {
	err := r.db.Debug().Create(&makanan).Error
	if err != nil {
		return err
	}

	return err
}

func NewMakananRepository(db *gorm.DB) IMakananRepository {
	return &MakananRepository{
		db: db,
	}
}
