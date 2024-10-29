package rest

import (
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

// penggunaHandler struct dengan field untuk mengakses usecase
type PenggunaHandler struct {
	Usecase *usecase.Usecase
}

// NewPenggunaHandler constructor untuk membuat instance penggunaHandler
func NewPenggunaHandler(usecase *usecase.Usecase) *PenggunaHandler {
	return &PenggunaHandler{
		Usecase: usecase,
	}
}

func (h *PenggunaHandler) PengunaRegister(c *gin.Context) {
	param := model.PengunaRegister{}

	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Usecase.PenggunaUsecase.RegisterPengguna(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "failed to register new user", err)
		return
	}

	response.Success(c, http.StatusCreated, "success register new user", nil)
}

// GetPengguna method handler untuk endpoint /tes
func (h *PenggunaHandler) GetAllPengguna(c *gin.Context) {
	
	allPengguna, err := h.Usecase.PenggunaUsecase.GetAllPengguna()
	if err != nil {
		response.Error(c, 404, "Failed to get all Pengguna", err)
	}
	response.Success(c, 200, "Hello world", allPengguna)
	
}