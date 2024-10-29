package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type ProfilPenggunaHandler struct {
    Usecase *usecase.Usecase
}

func NewprofilPenggunaHandler(usecase *usecase.Usecase) *ProfilPenggunaHandler{
    return &ProfilPenggunaHandler{
		Usecase: usecase,
	}
}
