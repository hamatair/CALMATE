package usecase

import (
	"errors"
	"math"
	"time"

	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/google/uuid"
)

type IPenggunaUsecase interface {
	// daftar fungsi
	LoginPengguna(model.PenggunaParam) (model.PenggunaLoginResponse, error)
	GetPengguna(model.PenggunaParam) (entity.Pengguna, error)
	GetAllPengguna() ([]entity.Pengguna, error)
	RegisterPengguna(model.PengunaRegister) error
}

type penggunaUsecase struct {
	Repository repository.Repository
	bcrypt     bcrypt.Interface
	jwtAuth    jwt.Interface
}

// LoginPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) LoginPengguna(param model.PenggunaParam) (model.PenggunaLoginResponse, error) {
	result := model.PenggunaLoginResponse{}

	email := model.PenggunaParam{
		Email: param.Email,
	}

	pengguna, err := u.Repository.PenggunaRepository.GetPengguna(email)
	if err != nil {
		return result, err
	}else if pengguna.IDPengguna == ""{
		return result, errors.New("email tidak ditemukan")
	}

	err = u.bcrypt.CompareAndHashPassword(pengguna.Password, param.Password)
	if err != nil {
    	return result, err
	}

	// Parsing IDPengguna dari string ke uuid.UUID
	idPengguna, err := uuid.Parse(pengguna.IDPengguna)
	if err != nil {
		return result, err
	}

// Membuat token JWT menggunakan UUID yang telah diparse
	token, err := u.jwtAuth.CreateJWTToken(idPengguna)
	if err != nil {
		return result, err
	}

	result.Token = token

	return result, nil
}

// GetPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) GetPengguna(param model.PenggunaParam) (entity.Pengguna, error) {
	Pengguna, err := u.Repository.PenggunaRepository.GetPengguna(param)
	if err != nil {
		return entity.Pengguna{}, err
	}
	return Pengguna, err
}

// RegisterPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) RegisterPengguna(param model.PengunaRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	pengguna := entity.Pengguna{
		IDPengguna: uuid.New().String(),
		Email:      param.Email,
		Password:   hashPassword,

		Role: "Pengguna",
	}

	err = u.Repository.PenggunaRepository.CreatePengguna(pengguna)
	if err != nil {
		return err
	}

	profilPengguna := entity.ProfilPengguna{
		IDProfil:          uuid.New().String(),
		IDPengguna:        pengguna.IDPengguna,
		NamaPengguna:      param.NamaPengguna,
		TanggalLahir: time.Now(),
		JenisKelamin:      param.JenisKelamin,
		TinggiBadan:       param.TinggiBadan,
		BeratBadan:        param.BeratBadan,
		Umur:              param.Umur,
		AktivitasPengguna: param.AktivitasPengguna,
	}

	err = u.Repository.ProfilPenggunaRepository.CreateProfilPengguna(profilPengguna)
	if err != nil {
		return err
	}

	riwayatKesehatan := entity.RiwayatKesehatan{
		IDRiwayat: uuid.New().String(),
		IDPengguna: pengguna.IDPengguna,
		
		NilaiBMI: profilPengguna.BeratBadan / float32(math.Pow(float64(profilPengguna.TinggiBadan) / 100, 2)),
	}

	err = u.Repository.RiwayatKesehatanRepository.CreateRiwayatKesehatan(riwayatKesehatan)
	if err != nil {
		return err
	}

	progres := entity.ProgresNutrisiHarian{
		IDProgresNutrisiHarian: uuid.New().String(),
		IDPengguna: pengguna.IDPengguna,
		JumlahKonsumsiKalori: 0,
		JumlahKonsumsiKarbohidrat: 0,
		JumlahKonsumsiProtein: 0,
		JumlahKonsumsiLemak: 0,
	}

	err = u.Repository.ProgresNutrisiHarian.CreateProgres(progres)
	if err != nil {
		return err
	}

	return err
}

// GetPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) GetAllPengguna() ([]entity.Pengguna, error) {
	allPengguna, err := u.Repository.PenggunaRepository.GetAllPengguna()
	if err != nil {
		return []entity.Pengguna{}, err
	}
	return allPengguna, err
}

func NewpenggunaUsecase(repository repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IPenggunaUsecase {
	return &penggunaUsecase{
		Repository: repository,
		jwtAuth:    jwtAuth,
		bcrypt:     bcrypt,
	}
}
