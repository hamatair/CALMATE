package usecase

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/google/uuid"
)

type IPenggunaUsecase interface {
	// daftar fungsi
	GetAllPengguna() ([]entity.Pengguna, error)
	RegisterPengguna(model.PengunaRegister) error
}

type penggunaUsecase struct {
	Repository repository.Repository
	bcrypt     bcrypt.Interface
	jwtAuth    jwt.Interface
}

// RegisterPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) RegisterPengguna(param model.PengunaRegister) error {
	hashPassword, err := u.bcrypt.GenerateFromPassword(param.Password)
	if err != nil {
		return err
	}

	NewIDPengguna := uuid.New()
	NewIDProfilPengguna := uuid.New()

	pengguna := entity.Pengguna{
		IDPengguna: NewIDPengguna.String(),
		Email:      param.Email,
		Password:   hashPassword,

		Role: "Pengguna",
	}

	profilPengguna := entity.ProfilPengguna{
		IDProfil:          NewIDProfilPengguna.String(),
		IDPengguna:        NewIDPengguna.String(),
		NamaPengguna:      param.NamaPengguna,
		JenisKelamin:      param.JenisKelamin,
		TinggiBadan:       param.TinggiBadan,
		BeratBadan:        param.BeratBadan,
		Umur:              param.Umur,
		AktivitasPengguna: param.AktivitasPengguna,
	}

	err = u.Repository.PenggunaRepository.CreatePengguna(pengguna)
	if err != nil {
		return err
	}

	err = u.Repository.ProfilPenggunaRepository.CreateProfilPengguna(profilPengguna)
	if err != nil {
		return err
	}

	return nil
}

// GetPengguna implements IPenggunaUsecase.
func (u *penggunaUsecase) GetAllPengguna() ([]entity.Pengguna, error) {
	allPengguna, err := u.Repository.PenggunaRepository.GetAllPengguna()
	if err != nil{
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
