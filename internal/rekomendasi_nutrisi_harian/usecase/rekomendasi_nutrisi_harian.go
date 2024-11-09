package usecase

import (
	"errors"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/google/uuid"
)

type IRekomendasiNutrisiHarianUsecase interface {
	GetRekomendasi(model.PenggunaParam) (entity.RekomendasiNutrisiHarian, error)
}

type rekomendasiNutrisiHarianUsecase struct {
	Repository *repository.Repository
}

// CreateRekomendasi implements IRekomendasiNutrisiHarianUsecase.
func (u *rekomendasiNutrisiHarianUsecase) GetRekomendasi(param model.PenggunaParam) (entity.RekomendasiNutrisiHarian, error) {
	profil, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return entity.RekomendasiNutrisiHarian{}, err
	}

	err = cekKelengkapanData(profil)
	if err != nil {
		return entity.RekomendasiNutrisiHarian{}, err
	}

	rekomendasi, err := u.Repository.RekomendasiNutrisiHarianRepository.GetRekomendasi(param) 
	if err != nil ||  isEmpty(rekomendasi){
		rekomendasi.IDRekomendasi = uuid.New().String()
		rekomendasi.IDPengguna = param.IDPengguna
		newRekomendasi, err := hitungRekomendasi(profil, rekomendasi)
		if err != nil {
			return entity.RekomendasiNutrisiHarian{}, err
		}
		
		err = u.Repository.RekomendasiNutrisiHarianRepository.CreateRekomendasi(newRekomendasi)
		if err != nil {
			return entity.RekomendasiNutrisiHarian{}, err
		}

		return newRekomendasi, err
	}

	newRekomendasi, err := hitungRekomendasi(profil, rekomendasi)
	if err != nil {
		return entity.RekomendasiNutrisiHarian{}, err
	}

	err = u.Repository.RekomendasiNutrisiHarianRepository.UpdateRekomendasi(param, newRekomendasi)

	return newRekomendasi, err
}

func NewrekomendasiNutrisiHarianUsecase(repository repository.Repository) IRekomendasiNutrisiHarianUsecase{
	return &rekomendasiNutrisiHarianUsecase{
		Repository: &repository,
	}
}

func hitungRekomendasi(profil entity.ProfilPengguna, rekomendasi entity.RekomendasiNutrisiHarian) (entity.RekomendasiNutrisiHarian, error) {
	if profil.JenisKelamin == "l"{
		rekomendasi.JumlahKaloriHarian = 
		66 + (13.7 * profil.BeratBadan) + (5 * profil.TinggiBadan) - (6.8 * float32(profil.Umur))
	} else if profil.JenisKelamin == "p" {
		rekomendasi.JumlahKaloriHarian = 
		655 + (9.6 * profil.BeratBadan) + (1.8 * profil.TinggiBadan) - (4.7 * float32(profil.Umur))
	}else {
		return rekomendasi, errors.New("invalid jenis kelamin type")
	}

	if profil.AktivitasPengguna == "jarang_olahraga" {
		rekomendasi.JumlahKaloriHarian *= 1.375 
	} else if profil.AktivitasPengguna == "cukup_olahraga" {
		rekomendasi.JumlahKaloriHarian *= 1.55
	} else if profil.AktivitasPengguna == "sering_olahraga" {
		rekomendasi.JumlahKaloriHarian *= 1.725
	} else {
		return rekomendasi, errors.New("invalid tktivitas type")
	}

	rekomendasi.AsupanKarbohidratHarian = (rekomendasi.JumlahKaloriHarian * 0.6) / 4
	rekomendasi.AsupanProteinHarian = (rekomendasi.JumlahKaloriHarian * 0.15) / 4
	rekomendasi.AsupanLemakHarian = (rekomendasi.JumlahKaloriHarian * 0.15) / 9

	return rekomendasi, nil
}

func cekKelengkapanData(profil entity.ProfilPengguna) error {
	if profil.JenisKelamin == "" || profil.Umur == 0 || profil.AktivitasPengguna == "" || profil.BeratBadan == 0 || profil.TinggiBadan == 0 {
        return errors.New("data profil tidak lengkap")
    }
	return nil
}

func isEmpty(profil entity.RekomendasiNutrisiHarian) bool {
    return profil == entity.RekomendasiNutrisiHarian{}
}