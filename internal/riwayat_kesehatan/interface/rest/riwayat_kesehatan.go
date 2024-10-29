package rest

import "github.com/bccfilkom-be/go-server/internal/usecase"

type RiwayatKesehatanHandler struct {
    Usecase *usecase.Usecase

}

func NewriwayatKesehatanHandler(usecase *usecase.Usecase) *RiwayatKesehatanHandler{
    return &RiwayatKesehatanHandler{
		Usecase: usecase,
	}
}
