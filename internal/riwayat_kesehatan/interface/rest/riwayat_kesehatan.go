package rest

import (
	"errors"
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type RiwayatKesehatanHandler struct {
    Usecase *usecase.Usecase

}

func NewriwayatKesehatanHandler(usecase *usecase.Usecase) *RiwayatKesehatanHandler{
    return &RiwayatKesehatanHandler{
		Usecase: usecase,
	}
}

func (h *RiwayatKesehatanHandler) GetRiwayatKesehatan(c *gin.Context){
	pengguna , ok := c.Get("pengguna")
	if !ok {
		response.Error(c, http.StatusNotFound, "Failed to get Pengguna", errors.New(""))
	}

	param , ok := pengguna.(model.PenggunaParam)
	if !ok {
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return	
	}

	riwayatKesehatan, err := h.Usecase.RiwayatKesehatanUsecase.GetRiwayatKesehatan(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get Riwayat Kesehatan", err)
	}

	response.Success(c, http.StatusOK, "Success to Get Riwayat Kesehatan", riwayatKesehatan)

}