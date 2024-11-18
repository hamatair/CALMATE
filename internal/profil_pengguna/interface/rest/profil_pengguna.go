package rest

import (
	"errors"
	"net/http"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/bccfilkom-be/go-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type ProfilPenggunaHandler struct {
    Usecase *usecase.Usecase
}

func NewprofilPenggunaHandler(usecase *usecase.Usecase) *ProfilPenggunaHandler{
    return &ProfilPenggunaHandler{
		Usecase: usecase,
	}
}

func (h *ProfilPenggunaHandler) GetProfilPengguna(c *gin.Context) {
	pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
		return
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}
	
	profilPengguna, err := h.Usecase.ProfilPenggunaUsecase.GetProfilPengguna(param)
	if err != nil {
		response.Error(c, 404, "Failed to get Profil Pengguna", err)
		return
	}
	response.Success(c, 200, "Success to Get Profil Pengguna", profilPengguna)
	
}

func (h *ProfilPenggunaHandler) UpdateProfilPengguna(c *gin.Context){
	pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
		return
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}

	newProfil := model.ProfilPengguna{}

	err := c.ShouldBindJSON(&newProfil)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Usecase.ProfilPenggunaUsecase.UpdateProfilPengguna(param, newProfil, model.Foto{Foto: nil}, false)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to Update Profil Pengguna", err)
		return
	}

	response.Success(c, http.StatusOK, "SUccess to Update Profil Pengguna", nil)
}

func (h *ProfilPenggunaHandler) DeleteFotoProfilPengguna(c *gin.Context) {
	pengguna, ok := c.Get("pengguna") 
	if !ok {
		response.Error(c, http.StatusBadRequest, "Failed to bind input", errors.New(""))
	}

	param, ok := pengguna.(model.PenggunaParam)
	if !ok {
		response.Error(c, http.StatusInternalServerError, "Failed to Casting", errors.New(""))
	}

	err := h.Usecase.ProfilPenggunaUsecase.DeleteFotoProfilPengguna(param)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to Delete Foto Profil Pengguna", err)
	}

	response.Success(c, http.StatusOK, "Success to delete foto profil Pengguna", err)
}

func (h *ProfilPenggunaHandler) UploadFotoProfilPengguna(c *gin.Context){
	pengguna, ok := c.Get("pengguna")
	if !ok {
		response.Error(c, 404, "Failed Get Login Pengguna", errors.New(""))
	}

	param, ok := pengguna.(model.PenggunaParam) 
	if !ok{
		response.Error(c, 500, "Failed to Cast Pengguna", errors.New("invalid user type"))
		return
	}

	foto, err := c.FormFile("foto")
	if err != nil {
		response.Success(c, http.StatusBadRequest, "failed to bind input", err)
		return
	}

	err = h.Usecase.ProfilPenggunaUsecase.UpdateProfilPengguna(param, model.ProfilPengguna{}, model.Foto{Foto: foto}, true)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to Update Profil Pengguna", err)
		return
	}

	response.Success(c, http.StatusOK, "Success to Update Profil Pengguna", nil)
}