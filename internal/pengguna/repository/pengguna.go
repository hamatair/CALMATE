package repository

import (
	"errors"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"gorm.io/gorm"
)

type IPenggunaRepository interface {
	GetPengguna(model.PenggunaParam) (entity.Pengguna, error)
	CreatePengguna(entity.Pengguna) error
	GetAllPengguna() ([]entity.Pengguna, error)
}

type PenggunaRepository struct {
	db *gorm.DB
}

// GetPengguna implements IPenggunaRepository.
func (r *PenggunaRepository) GetPengguna(param model.PenggunaParam) (entity.Pengguna, error) {
	var Pengguna entity.Pengguna
	err := r.db.Debug().Where(&param).Find(&Pengguna).Error
	if err != nil {
		return entity.Pengguna{}, err
	}

	return Pengguna, nil}

// GetAllPengguna implements IPenggunaRepository.
func (r *PenggunaRepository) GetAllPengguna() ([]entity.Pengguna, error) {
	var allPengguna []entity.Pengguna // Menggunakan slice untuk menampung banyak data
	err := r.db.Debug().Find(&allPengguna).Error
	if err != nil {
		return []entity.Pengguna{}, err
	}

	return allPengguna, nil

}

// CreatePengguna implements IPenggunaRepository.
func (r *PenggunaRepository) CreatePengguna(param entity.Pengguna) error {
	var existingUser entity.Pengguna
	err := r.db.Where("email = ?", param.Email).First(&existingUser).Error
	if err == nil {
		return errors.New("email Sudah Ada")
	}
	
	err = r.db.Debug().Create(&param).Error
	if err != nil {
		return err
	}

	return nil
}

func NewPenggunaRepository(db *gorm.DB) IPenggunaRepository {
	return &PenggunaRepository{
		db: db,
	}
}
