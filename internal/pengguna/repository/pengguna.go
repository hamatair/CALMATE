package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"gorm.io/gorm"
)

type IPenggunaRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
	CreatePengguna(entity.Pengguna) error
	GetAllPengguna() ([]entity.Pengguna, error)
}

type PenggunaRepository struct {
	db *gorm.DB
}

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
	err := r.db.Debug().Create(&param).Error
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
