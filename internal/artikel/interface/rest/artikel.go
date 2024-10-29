package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type ArtikelHandler struct {
    Usecase *usecase.Usecase
}

func NewartikelHandler(usecase *usecase.Usecase) *ArtikelHandler{
    return &ArtikelHandler{
        Usecase: usecase,
    }
}
