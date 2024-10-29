package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"gorm.io/gorm"
)

type IProfilPenggunaRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
	CreateProfilPengguna(entity.ProfilPengguna) error
}

type ProfilPenggunaRepository struct {
	db *gorm.DB
}

// CreateProfilPengguna implements IProfilPenggunaRepository.
func (r *ProfilPenggunaRepository) CreateProfilPengguna(param entity.ProfilPengguna) error {
	err := r.db.Debug().Create(&param).Error
	if err != nil {
		return err
	}

	return nil
}

func NewProfilPenggunaRepository(db *gorm.DB) IProfilPenggunaRepository {
	return &ProfilPenggunaRepository{
		db: db,
	}
}
