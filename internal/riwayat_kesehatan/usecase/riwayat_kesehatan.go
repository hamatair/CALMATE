package usecase

import (
	"errors"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IRiwayatKesehatanUsecase interface {
	GetRiwayatKesehatan(model.PenggunaParam) (entity.RiwayatKesehatan, error)
}

type riwayatKesehatanUsecase struct {
	Repository repository.Repository}

// GetRiwayatKesehatan implements IRiwayatKesehatanUsecase.
func (u *riwayatKesehatanUsecase) GetRiwayatKesehatan(param model.PenggunaParam) (entity.RiwayatKesehatan, error) {
	riwayatKesehatan, err := u.Repository.RiwayatKesehatanRepository.GetRiwayatKesehatan(param)
	if err != nil {
		return riwayatKesehatan, err
	}

    if riwayatKesehatan.IDPengguna == "" {
		return entity.RiwayatKesehatan{}, errors.New("profil pengguna tidak ditemukan")
	}

	return riwayatKesehatan, err
}

func NewriwayatKesehatanUsecase(repository repository.Repository) IRiwayatKesehatanUsecase {
	return &riwayatKesehatanUsecase{
		Repository: repository,
	}
}
