package rest

import (
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type MakananHandler struct {
    Usecase *usecase.Usecase
}

func NewmakananHandler(usecase *usecase.Usecase) *MakananHandler{
    return &MakananHandler{
        Usecase: usecase,
    }
}

func (h *MakananHandler) CreateMakanan(c *gin.Context) {
    param := model.Makanan{}

    err := c.ShouldBindJSON(&param)
    if err != nil {
        response.Error(c, http.StatusBadRequest, "Failed to bind JSON", err)
        return
    }

    err = h.Usecase.Makanan.CreateMakanan(param)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to Create Makanan", err)
        return
    }

    response.Success(c, http.StatusOK, "Success to Create Makanan", err)

} 

func (h *MakananHandler) GetMakanan(c *gin.Context){
    param := model.GetMakanan{}
    err := c.ShouldBindJSON(&param)
    if err != nil {
        response.Error(c, http.StatusBadRequest, "Failed to Bind JSON", err)
        return
    }

    makanan, err := h.Usecase.Makanan.GetMakanan(param.Nama)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to Get Makanan", err)
    }

    response.Success(c, http.StatusOK, "Success to Get Makanan", makanan)
}