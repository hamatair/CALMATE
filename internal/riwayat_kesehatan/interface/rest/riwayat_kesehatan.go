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

func (h *RiwayatKesehatanHandler) UpdateRiwayatKesehatan(c *gin.Context){
	pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}

	newRiwayatKesehatan := model.UpdateRiwayatKesehatan{}

	err := c.ShouldBindJSON(&newRiwayatKesehatan)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Usecase.RiwayatKesehatanUsecase.UpdateRiwayatKesehatan(param, newRiwayatKesehatan)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to Update Riwayat Kesehatan Pengguna", err)
		return
	}

	response.Success(c, http.StatusOK, "SUccess to Update Riwayat Kesehatan Pengguna", nil)
}

func (h *RiwayatKesehatanHandler) DeleteRiwayatKesehatan (c *gin.Context) {
	pengguna, ok := c.Get("pengguna") 
	if !ok {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", errors.New(""))
	}

	del := model.DeleteRiwayatKesehatan{}

	err := c.ShouldBindJSON(&del)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	param, ok := pengguna.(model.PenggunaParam)
	if !ok {
		response.Error(c, http.StatusInternalServerError, "Failed to Casting", errors.New(""))
		return
	}

	err = h.Usecase.RiwayatKesehatanUsecase.DeleteRiwayatKesehatan(param, del)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to Delete Riwayat Kesehatan", err)
		return
	}

	response.Success(c, http.StatusOK, "Success to delete Riwayat Kesehatan Pengguna", err)

}