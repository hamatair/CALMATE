package repository

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"gorm.io/gorm"
)

type IRiwayatKesehatanRepository interface {
	// CreateUser(user entity.User) (entity.User, error)
	// GetUser(param model.UserParam) (entity.User, error)
	CreateRiwayatKesehatan(entity.RiwayatKesehatan) error
	GetRiwayatKesehatan(model.PenggunaParam) (entity.RiwayatKesehatan, error)
	UpdateRiwayatKesehatan(model.PenggunaParam, entity.RiwayatKesehatan) error
	DeleteRiwayatKesehatan(model.PenggunaParam, entity.RiwayatKesehatan) error
}

type RiwayatKesehatanRepository struct {
	db *gorm.DB
}

// DeleteRiwayatKesehatan implements IRiwayatKesehatanRepository.
func (r *RiwayatKesehatanRepository) DeleteRiwayatKesehatan(param model.PenggunaParam, data entity.RiwayatKesehatan) error {
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).Save(data).Error
	return err
}

// UpdateRiwayatKesehatan implements IRiwayatKesehatanRepository.
func (r *RiwayatKesehatanRepository) UpdateRiwayatKesehatan(param model.PenggunaParam, newRiwayat entity.RiwayatKesehatan) error {
	err := r.db.Debug().Where("id_pengguna = ?", param.IDPengguna).Save(newRiwayat).Error
	return err
}

// GetRiwayatKesehatan implements IRiwayatKesehatanRepository.
func (r *RiwayatKesehatanRepository) GetRiwayatKesehatan(param model.PenggunaParam) (entity.RiwayatKesehatan, error) {
	var riwayatKesehatan entity.RiwayatKesehatan
	err := r.db.Debug().Where("id_pengguna = ?", &param.IDPengguna).First(&riwayatKesehatan).Error
	if err != nil {
		return entity.RiwayatKesehatan{}, err
	}

	return riwayatKesehatan, err
}

// CreateRiwayatKesehatan implements IRiwayatKesehatanRepository.
func (r *RiwayatKesehatanRepository) CreateRiwayatKesehatan(param entity.RiwayatKesehatan) error {
	err := r.db.Debug().Create(&param).Error
	if err != nil {
		return err
	}

	return nil
}

func NewRiwayatKesehatanRepository(db *gorm.DB) IRiwayatKesehatanRepository {
	return &RiwayatKesehatanRepository{
		db: db,
	}
}
