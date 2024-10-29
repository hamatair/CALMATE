package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type RekomendasiNutrisiHarianHandler struct {
    Usecase *usecase.Usecase
}

func NewrekomendasiNutrisiHarianHandler(usecase *usecase.Usecase) *RekomendasiNutrisiHarianHandler {
    return &RekomendasiNutrisiHarianHandler{
        Usecase: usecase,
    }
}
