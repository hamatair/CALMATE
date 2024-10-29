package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type ProgresNutrisiHarianHandler struct {
    Usecase *usecase.Usecase
}

func NewprogresNutrisiHarianHandler(usecase *usecase.Usecase) *ProgresNutrisiHarianHandler{
    return &ProgresNutrisiHarianHandler{
        Usecase: usecase,
    }
}