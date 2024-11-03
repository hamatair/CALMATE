package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"gorm.io/gorm"
)

type IProfilPenggunaRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
	CreateProfilPengguna(entity.ProfilPengguna) error
	GetProfilPengguna(model.PenggunaParam) (entity.ProfilPengguna, error)
	UpdateProfilPengguna(model.PenggunaParam, model.ProfilPengguna) error
	DeleteFotoProfilPengguna(model.PenggunaParam) error
}

type ProfilPenggunaRepository struct {
	db *gorm.DB
}

// DeleteFotoProfilPengguna implements IProfilPenggunaRepository.
func (r *ProfilPenggunaRepository) DeleteFotoProfilPengguna(param model.PenggunaParam) error {
	err := r.db.Debug().
		Model(&model.ProfilPengguna{}).
		Where("id_pengguna = ?", param.IDPengguna).
		Update("foto", nil).Error
	if err != nil {
		return err
	}
	return err
}

// UpdateProfilPengguna implements IProfilPenggunaRepository.
func (r *ProfilPenggunaRepository) UpdateProfilPengguna(param model.PenggunaParam, newProfil model.ProfilPengguna) error {
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).Save(newProfil).Error
	return err
}

// GetProfilPengguna implements IProfilPenggunaRepository.
func (r *ProfilPenggunaRepository) GetProfilPengguna(param model.PenggunaParam) (entity.ProfilPengguna, error) {
	var profilPengguna entity.ProfilPengguna
	err := r.db.Debug().Where("id_pengguna = ?", &param.IDPengguna).Find(&profilPengguna).Error
	if err != nil {
		return entity.ProfilPengguna{}, err
	}

	return profilPengguna, err
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
