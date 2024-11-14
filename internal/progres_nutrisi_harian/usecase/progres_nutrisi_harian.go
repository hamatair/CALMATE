package usecase

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IProgresNutrisiHarianUsecase interface {
	GetProges(model.PenggunaParam) (entity.ProgresNutrisiHarian, error)
	UpdateProgres(model.PenggunaParam, model.ProgresNutrisiHarian) error
	ResetAllProgres() error
}

type progresNutrisiHarianUsecase struct {
	Repository repository.Repository
}

// ResetProgres implements IProgresNutrisiHarianUsecase.
func (u *progresNutrisiHarianUsecase) ResetAllProgres() error {
	AllProgres, err := u.Repository.ProgresNutrisiHarian.GetAllProgres()
	if err != nil {
		return err
	}

	for i:= range AllProgres{
		AllProgres[i].JumlahKonsumsiKalori = 0
		AllProgres[i].JumlahKonsumsiKarbohidrat = 0
		AllProgres[i].JumlahKonsumsiProtein = 0
		AllProgres[i].JumlahKonsumsiLemak = 0
	}

	err = u.Repository.ProgresNutrisiHarian.ResetAllProgres(AllProgres)
	if err != nil {
		return err
	}

	return err

}

// GetProges implements IProgresNutrisiHarianUsecase.
func (u *progresNutrisiHarianUsecase) GetProges(param model.PenggunaParam) (entity.ProgresNutrisiHarian, error) {
	progres, err := u.Repository.ProgresNutrisiHarian.GetProgres(param)
	if err != nil {
		return entity.ProgresNutrisiHarian{}, err
	}

	return progres, err
}

// UpdateProgres implements IProgresNutrisiHarianUsecase.
func (u *progresNutrisiHarianUsecase) UpdateProgres(param model.PenggunaParam, progres model.ProgresNutrisiHarian) error {
	oldProgres, err := u.Repository.ProgresNutrisiHarian.GetProgres(param)
	if err != nil {
		return err
	}

	newProgres := model.ProgresNutrisiHarian{
		JumlahKonsumsiKalori:      oldProgres.JumlahKonsumsiKalori + progres.JumlahKonsumsiKalori,
		JumlahKonsumsiKarbohidrat: oldProgres.JumlahKonsumsiKarbohidrat + progres.JumlahKonsumsiKarbohidrat,
		JumlahKonsumsiProtein:     oldProgres.JumlahKonsumsiProtein + progres.JumlahKonsumsiProtein,
		JumlahKonsumsiLemak:       oldProgres.JumlahKonsumsiLemak + progres.JumlahKonsumsiLemak,
	}

	err = u.Repository.ProgresNutrisiHarian.UpdateProgres(param, newProgres)
	if err != nil {
		return err
	}

	return err
}

func NewprogresNutrisiHarianUsecase(repository repository.Repository) IProgresNutrisiHarianUsecase {
	return &progresNutrisiHarianUsecase{
		Repository: repository,
	}
}
