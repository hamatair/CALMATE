package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"gorm.io/gorm"
)

type IProgresNutrisiHarianRepository interface {
	CreateProgres(entity.ProgresNutrisiHarian) error
	UpdateProgres(model.PenggunaParam, model.ProgresNutrisiHarian) error
	GetProgres(model.PenggunaParam) (entity.ProgresNutrisiHarian, error)
	GetAllProgres() ([]entity.ProgresNutrisiHarian, error)
}

type ProgresNutrisiHarianRepository struct {
	db *gorm.DB
}

// GetAllProgres implements IProgresNutrisiHarianRepository.
func (r *ProgresNutrisiHarianRepository) GetAllProgres() ([]entity.ProgresNutrisiHarian, error) {
	var AllProgres []entity.ProgresNutrisiHarian
	err := r.db.Debug().Find(&AllProgres).Error
	if err != nil {
		return []entity.ProgresNutrisiHarian{}, err
	}

	return AllProgres, err
}

// CreateProgres implements IProgresNutrisiHarianRepository.
func (r *ProgresNutrisiHarianRepository) CreateProgres(progres entity.ProgresNutrisiHarian) error {
	err := r.db.Debug().Create(&progres).Error
	if err != nil {
		return err
	}

	return err
}

// GetProgres implements IProgresNutrisiHarianRepository.
func (r *ProgresNutrisiHarianRepository) GetProgres(param model.PenggunaParam) (entity.ProgresNutrisiHarian, error) {
	var progres entity.ProgresNutrisiHarian
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).First(&progres).Error
	if err != nil {
		return entity.ProgresNutrisiHarian{}, err
	}
	return progres, err
}

// UpdateProgres implements IProgresNutrisiHarianRepository.
func (r *ProgresNutrisiHarianRepository) UpdateProgres(param model.PenggunaParam, newProgres model.ProgresNutrisiHarian) error {
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).Save(newProgres).Error
	if err != nil {
		return err
	}

	return err
}

func NewProgresNutrisiHarianRepository(db *gorm.DB) IProgresNutrisiHarianRepository {
	return &ProgresNutrisiHarianRepository{
		db: db,
	}
}
