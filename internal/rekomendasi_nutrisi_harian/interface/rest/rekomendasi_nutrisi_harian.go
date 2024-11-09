package rest

import (
	"errors"
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type RekomendasiNutrisiHarianHandler struct {
    Usecase *usecase.Usecase
}

func NewrekomendasiNutrisiHarianHandler(usecase *usecase.Usecase) *RekomendasiNutrisiHarianHandler {
    return &RekomendasiNutrisiHarianHandler{
        Usecase: usecase,
    }
}

func (h *RekomendasiNutrisiHarianHandler) GetRekomendasi(c *gin.Context){
    param, ok := c.Get("pengguna")
    if !ok {
        response.Error(c, http.StatusBadRequest, "failed to get pengguna", errors.New(""))
        return
    }

    pengguna := param.(model.PenggunaParam)

    rekomendasi, err := h.Usecase.RekomendasiNutrisiHarianUsecase.GetRekomendasi(pengguna)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "failed to Get Rekomendasi", err)
        return
    }

    response.Success(c, http.StatusOK, "Success to Get rekomendasi", rekomendasi)
}
