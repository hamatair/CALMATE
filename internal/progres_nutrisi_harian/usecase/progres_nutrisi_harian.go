package usecase

import "github.com/bccfilkom-be/go-server/internal/progres_nutrisi_harian/repository"

type IProgresNutrisiHarianUsecase interface {}

type progresNutrisiHarianUsecase struct {
    ProgresNutrisiHarianRepository repository.IProgresNutrisiHarianRepository
}

func NewprogresNutrisiHarianUsecase(progresNutrisiHarianRepository repository.IProgresNutrisiHarianRepository) IProgresNutrisiHarianUsecase {
    return &progresNutrisiHarianUsecase{
        ProgresNutrisiHarianRepository: progresNutrisiHarianRepository,
    }
}
