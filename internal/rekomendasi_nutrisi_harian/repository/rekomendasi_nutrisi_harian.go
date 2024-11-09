package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"gorm.io/gorm"
)

type IRekomendasiNutrisiharianRepository interface {
	CreateRekomendasi(entity.RekomendasiNutrisiHarian) error
	GetRekomendasi(model.PenggunaParam) (entity.RekomendasiNutrisiHarian, error)
	UpdateRekomendasi(model.PenggunaParam, entity.RekomendasiNutrisiHarian) error
}

type RekomendasiNutrisiHarianRepository struct {
	db *gorm.DB
}

// UpdateRekomendasi implements IRekomendasiNutrisiharianRepository.
func (r *RekomendasiNutrisiHarianRepository) UpdateRekomendasi(param model.PenggunaParam, newRekomendasi entity.RekomendasiNutrisiHarian) error {
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).Save(newRekomendasi).Error
	return err
}

// GetRekomendasi implements IRekomendasiNutrisiharianRepository.
func (r *RekomendasiNutrisiHarianRepository) GetRekomendasi(param model.PenggunaParam) (entity.RekomendasiNutrisiHarian, error) {
	var rekomendasi entity.RekomendasiNutrisiHarian
	err := r.db.Debug().Where("id_pengguna = ?", &param.IDPengguna).Find(&rekomendasi).Error
	if err != nil {
		return entity.RekomendasiNutrisiHarian{}, err
	}

	return rekomendasi, err
}

// UpdateRekomendasi implements IRekomendasiNutrisiharianRepository.
func (r *RekomendasiNutrisiHarianRepository) CreateRekomendasi(rekomendasi entity.RekomendasiNutrisiHarian) error {
	err := r.db.Debug().Create(&rekomendasi).Error
	if err != nil {
		return err
	}

	return nil
}

func NewRekomendasiNutrisiHarianRepository(db *gorm.DB) IRekomendasiNutrisiharianRepository {
	return &RekomendasiNutrisiHarianRepository{
		db: db,
	}
}
