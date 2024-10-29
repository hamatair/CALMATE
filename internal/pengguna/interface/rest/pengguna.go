package rest

import (

	"github.com/bccfilkom-be/go-server/internal/usecase"
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

// GetPengguna method handler untuk endpoint /tes
func (h *PenggunaHandler) GetPengguna(c *gin.Context) {
	// Logic untuk mendapatkan data pengguna, bisa menggunakan usecase
	// Misalnya, kita panggil `h.usecase.SomeFunction()` untuk mengambil data

	// Response dummy untuk contoh

	response.Success(c, 200, "Hello world", nil)
}
