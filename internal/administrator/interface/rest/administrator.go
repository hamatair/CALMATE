package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type AdministratorHandler struct {
    Usecase *usecase.Usecase
}

func NewadministratorHandler(usecase *usecase.Usecase) *AdministratorHandler{
    return &AdministratorHandler{
        Usecase: usecase,
    }
}
