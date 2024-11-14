package rest

import (
	"errors"
	"log"
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProgresNutrisiHarianHandler struct {
    Usecase *usecase.Usecase
}

func NewprogresNutrisiHarianHandler(usecase *usecase.Usecase) *ProgresNutrisiHarianHandler{
    return &ProgresNutrisiHarianHandler{
        Usecase: usecase,
    }
}

func (h *ProgresNutrisiHarianHandler) UpdateProgres(c *gin.Context){
    pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}
    
    newProgres := model.ProgresNutrisiHarian{}

    err := c.ShouldBindJSON(&newProgres)
    if err != nil {
        response.Error(c, http.StatusBadRequest, "Failed to bind JSON", err)
        return
    }

    err = h.Usecase.ProgresNutrisiHarianUsecase.UpdateProgres(param, newProgres)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to Update Progres", err)
        return
    }

    response.Success(c, http.StatusOK, "Success to Update Progres", nil)
}

func (h *ProgresNutrisiHarianHandler) GetProgres(c *gin.Context){
    pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}

    progres, err := h.Usecase.ProgresNutrisiHarianUsecase.GetProges(param)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to Get Progres", err)
        return
    }

    response.Success(c, http.StatusOK, "Success to Get Progres", progres)
}

func (h *ProgresNutrisiHarianHandler) ResetProgres(){
    err := h.Usecase.ProgresNutrisiHarianUsecase.ResetAllProgres()
    if err != nil {
        log.Fatal(err)
    }
}