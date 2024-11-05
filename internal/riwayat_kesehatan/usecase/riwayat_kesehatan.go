package usecase

import (
	"errors"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/google/uuid"
)

type IRiwayatKesehatanUsecase interface {
	GetRiwayatKesehatan(model.PenggunaParam) (entity.RiwayatKesehatan, error)
	UpdateRiwayatKesehatan(model.PenggunaParam, model.UpdateRiwayatKesehatan) error
	DeleteRiwayatKesehatan(model.PenggunaParam, model.DeleteRiwayatKesehatan) error
}

type riwayatKesehatanUsecase struct {
	Repository repository.Repository
}

// DeleteRiwayatKesehatan implements IRiwayatKesehatanUsecase.
func (u *riwayatKesehatanUsecase) DeleteRiwayatKesehatan(param model.PenggunaParam, del model.DeleteRiwayatKesehatan) error {
	_, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return err
	}

	data, err := u.Repository.RiwayatKesehatanRepository.GetRiwayatKesehatan(param)
	if err != nil {
		return err
	}

	switch del.Jenis{
	case "alergi" :
		data.Alergi = append(data.Alergi[:del.Index], data.Alergi[del.Index+1:]...)
	case "riwayat_obat" :
		data.RiwayatObat = append(data.RiwayatObat[:del.Index], data.RiwayatObat[del.Index+1:]...)
	case "riwayat_operasi" :
		data.RiwayatOperasi = append(data.RiwayatOperasi[:del.Index], data.RiwayatOperasi[del.Index+1:]...)
	case "riwayat_penyakit" :
		data.RiwayatPenyakit = append(data.RiwayatPenyakit[:del.Index], data.RiwayatPenyakit[del.Index+1:]...)
	default:
		return errors.New("kategori tidak valid")
	}
	
	err = u.Repository.RiwayatKesehatanRepository.DeleteRiwayatKesehatan(param, data)
	if err != nil {
		return err
	}

	return err
}

// UpdateRiwayatKesehatan implements IRiwayatKesehatanUsecase.
func (u *riwayatKesehatanUsecase) UpdateRiwayatKesehatan(param model.PenggunaParam, newRiwayatKesehatan model.UpdateRiwayatKesehatan) error {
	_, err := u.Repository.ProfilPenggunaRepository.GetProfilPengguna(param)
	if err != nil {
		return err
	}

	data, err := u.Repository.RiwayatKesehatanRepository.GetRiwayatKesehatan(param)
	if err != nil {
		return err
	}

	newData := entity.DetailRiwayatKesehatan{
		ID: uuid.New().String(),
		Detail: newRiwayatKesehatan.Detail,
		Tanggal: newRiwayatKesehatan.Tanggal,
	}

	switch newRiwayatKesehatan.Jenis{
	case "alergi" :
		data.Alergi = append(data.Alergi, newData)
	case "riwayat_obat" :
		data.RiwayatObat = append(data.RiwayatObat, newData)
	case "riwayat_operasi" :
		data.RiwayatOperasi = append(data.RiwayatOperasi, newData)
	case "riwayat_penyakit" :
		data.RiwayatPenyakit = append(data.RiwayatPenyakit, newData)
	default:
		return errors.New("kategori tidak valid")
	}


	err = u.Repository.RiwayatKesehatanRepository.UpdateRiwayatKesehatan(param, data)
	if err != nil {
		return err
	}

	return err
}

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
