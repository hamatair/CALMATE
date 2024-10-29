package usecase

import "github.com/bccfilkom-be/go-server/internal/rekomendasi_nutrisi_harian/repository"

type IRekomendasiNutrisiHarianUsecase interface {}

type rekomendasiNutrisiHarianUsecase struct {
    RekomendasiNutrisiHarianRepository repository.IRekomendasiNutrisiharianRepository
}

func NewrekomendasiNutrisiHarianUsecase(rekomendasiNutrisiHarianRepository repository.IRekomendasiNutrisiharianRepository) IRekomendasiNutrisiHarianUsecase {
    return &rekomendasiNutrisiHarianUsecase{
        RekomendasiNutrisiHarianRepository: rekomendasiNutrisiHarianRepository,
    }
}
