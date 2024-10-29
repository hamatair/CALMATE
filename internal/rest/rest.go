package rest

import (
	"fmt"
	"log"
	"os"

	administrator "github.com/bccfilkom-be/go-server/internal/administrator/interface/rest"
	artikel "github.com/bccfilkom-be/go-server/internal/artikel/interface/rest"
	pengguna "github.com/bccfilkom-be/go-server/internal/pengguna/interface/rest"
	profilPengguna "github.com/bccfilkom-be/go-server/internal/profil_pengguna/interface/rest"
	progresNutrisiHarian "github.com/bccfilkom-be/go-server/internal/progres_nutrisi_harian/interface/rest"
	rekomendasiNutrisiHarian "github.com/bccfilkom-be/go-server/internal/rekomendasi_nutrisi_harian/interface/rest"
	riwayatKesehatan "github.com/bccfilkom-be/go-server/internal/riwayat_kesehatan/interface/rest"
	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Rest struct {
	router         *gin.Engine
	Usecase        *usecase.Usecase
	penggunaHandler *pengguna.PenggunaHandler
	profilPengguna *profilPengguna.ProfilPenggunaHandler
	riwayatKesehatan *riwayatKesehatan.RiwayatKesehatanHandler
	rekomendasiNutrisiHarian *rekomendasiNutrisiHarian.RekomendasiNutrisiHarianHandler
	progresNutrisiHarian *progresNutrisiHarian.ProgresNutrisiHarianHandler
	artikel *artikel.ArtikelHandler
	administrator *administrator.AdministratorHandler

}

// NewRest constructor untuk Rest, menginisialisasi penggunaHandler juga
func NewRest(usecase *usecase.Usecase) *Rest {
	return &Rest{
		router:         gin.Default(),
		Usecase:        usecase,
		penggunaHandler: pengguna.NewPenggunaHandler(usecase),
	}
}

func (r *Rest) MountEndpoint() {
	routerGroup := r.router.Group("/api/v1")

	// Menggunakan handler dari penggunaHandler
	routerGroup.GET("/tes", r.penggunaHandler.GetPengguna)
}

func (r *Rest) Serve() {
	addr := os.Getenv("APP_ADDRESS")
	port := os.Getenv("APP_PORT")

	if addr == "" {
		addr = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}

	err := r.router.Run(fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
