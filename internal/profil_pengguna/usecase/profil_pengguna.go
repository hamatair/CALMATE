package usecase

import (
	"errors"

	"github.com/bccfilkom-be/go-server/internal/repository"
	entity "github.com/bccfilkom-be/go-server/internal/domain"

	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IProfilPenggunaUsecase interface {
	GetProfilPengguna(model.PenggunaParam) (entity.ProfilPengguna, error)
	UpdateProfilPengguna(model.PenggunaParam, model.ProfilPengguna) error
	DeleteProfilPenggguna(model.PenggunaParam) error
}

type profilPenggunaUsecase struct {
	Repository repository.Repository
}

// DeleteProfilPenggguna implements IProfilPenggunaUsecase.
func (u *profilPenggunaUsecase) DeleteProfilPenggguna(param model.PenggunaParam) error {
	err := u.Repository.ProfilPenggunaRepository.DeleteFotoProfilPengguna(param)
	if err != nil {
		return err
	}

	return err
}

// UpdateProfilPengguna implements IProfilPenggunaUsecase.
func (u *profilPenggunaUsecase) UpdateProfilPengguna(param model.PenggunaParam, newProfil model.ProfilPengguna) error {
	_, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return err
	}

	err = u.Repository.ProfilPenggunaRepository.UpdateProfilPengguna(param, newProfil)
	if err != nil {
		return err
	}

	return err
}

// GetProfilPengguna implements IProfilPenggunaUsecase.
func (u *profilPenggunaUsecase) GetProfilPengguna(param model.PenggunaParam) (entity.ProfilPengguna, error) {
	profilPengguna, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return entity.ProfilPengguna{}, err
	}

	if profilPengguna.IDPengguna == "" {
		return entity.ProfilPengguna{}, errors.New("profil pengguna tidak ditemukan") // Mengembalikan error jika tidak ada record ditemukan
	}

	return profilPengguna, err
}

func NewProfilPenggunaUsecase(repository repository.Repository) IProfilPenggunaUsecase {
	return &profilPenggunaUsecase{
		Repository: repository,
	}
}
