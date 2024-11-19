package usecase

import (
	"errors"
	"fmt"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"

	"github.com/bccfilkom-be/go-server/pkg/database/supabase"
	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IProfilPenggunaUsecase interface {
	GetProfilPengguna(model.PenggunaParam) (entity.ProfilPengguna, error)
	UpdateProfilPengguna(model.PenggunaParam, model.ProfilPengguna, model.Foto, bool) error
	DeleteFotoProfilPengguna(model.PenggunaParam) error
}

type profilPenggunaUsecase struct {
	Repository repository.Repository
	Supabase   supabase.Interface
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
func (u *profilPenggunaUsecase) UpdateProfilPengguna(param model.PenggunaParam, newProfil model.ProfilPengguna, foto model.Foto, isFoto bool) error {
	oldProfil, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return err
	}

	var namaFoto string
	var fotoLink string

	if isFoto {
		// Jika ada foto lama, hapus file tersebut
			if oldProfil.LinkFoto != "" {
				folder := oldProfil.IDProfil + "/"
				file := oldProfil.NamaFoto
				fmt.Printf("Mencoba menghapus file: %s/%s\n", folder, file)
				err = u.Supabase.Delete(folder, file)
				if err != nil {
					return fmt.Errorf("gagal menghapus file lama: %v", err)
				}
			}


		// Upload foto baru
		namaFoto = foto.Foto.Filename
		fotoLink, err = u.Supabase.Upload(foto.Foto, oldProfil.IDProfil) // Gunakan IDProfil sebagai folder
		if err != nil {
			return err
		}

		// Update profil pengguna
		newProfil.NamaPengguna = oldProfil.NamaPengguna
		newProfil.TanggalLahir = oldProfil.TanggalLahir
		newProfil.JenisKelamin = oldProfil.JenisKelamin
		newProfil.TinggiBadan = oldProfil.TinggiBadan
		newProfil.BeratBadan = oldProfil.BeratBadan
		newProfil.Umur = oldProfil.Umur
		newProfil.AktivitasPengguna = oldProfil.AktivitasPengguna
		newProfil.Alamat = oldProfil.Alamat
		newProfil.NoTeleponPengguna = oldProfil.NoTeleponPengguna
		newProfil.NamaFoto = namaFoto
		newProfil.LinkFoto = fotoLink
	} else {
		// Jika tidak ada foto baru, gunakan foto lama
		newProfil.NamaFoto = oldProfil.NamaFoto
		newProfil.LinkFoto = oldProfil.LinkFoto
	}

	// Update profil pengguna di database
	err = u.Repository.ProfilPenggunaRepository.UpdateProfilPengguna(param, newProfil)
	if err != nil {
		return err
	}

	return nil
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
		Supabase:   supabase,
	}
}
