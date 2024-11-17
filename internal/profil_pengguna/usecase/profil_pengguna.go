package usecase

import (
	"errors"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"

	"github.com/bccfilkom-be/go-server/pkg/database/supabase"
	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IProfilPenggunaUsecase interface {
	GetProfilPengguna(model.PenggunaParam) (entity.ProfilPengguna, error)
	UpdateProfilPengguna(model.PenggunaParam, model.ProfilPengguna, model.Foto) error
	DeleteFotoProfilPengguna(model.PenggunaParam) error
}

type profilPenggunaUsecase struct {
	Repository repository.Repository
	Supabase supabase.Interface
}


// DeleteProfilPenggguna implements IProfilPenggunaUsecase.
func (u *profilPenggunaUsecase) DeleteFotoProfilPengguna(param model.PenggunaParam) error {
	err := u.Repository.ProfilPenggunaRepository.DeleteFotoProfilPengguna(param)
	if err != nil {
		return err
	}

	return err
}

// UpdateProfilPengguna implements IProfilPenggunaUsecase.
func (u *profilPenggunaUsecase) UpdateProfilPengguna(param model.PenggunaParam, newProfil model.ProfilPengguna, foto model.Foto) error {
	oldProfil, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return err
	}

	if oldProfil.LinkFoto != ""{
		err = u.Supabase.Delete([]string{oldProfil.NamaFoto})
		if err != nil {
			return err
		}
	}

	var link string

	if foto.Foto != nil {
		link, err = u.Supabase.Upload(foto.Foto)
		if err != nil {
			return err
		}
	}

	newProfil.NamaPengguna = oldProfil.NamaPengguna
	newProfil.TanggalLahir = oldProfil.TanggalLahir
	newProfil.JenisKelamin = oldProfil.JenisKelamin
    newProfil.TinggiBadan = oldProfil.TinggiBadan
    newProfil.BeratBadan = oldProfil.BeratBadan
    newProfil.Umur = oldProfil.Umur
    newProfil.AktivitasPengguna = oldProfil.AktivitasPengguna
    newProfil.Alamat = oldProfil.Alamat
    newProfil.NoTeleponPengguna = oldProfil.NoTeleponPengguna
    newProfil.NamaFoto  = foto.Foto.Filename
	newProfil.LinkFoto = link

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

func NewProfilPenggunaUsecase(repository repository.Repository, supabase supabase.Interface) IProfilPenggunaUsecase {
	return &profilPenggunaUsecase{
		Repository: repository,
		Supabase: supabase,
	}
}
