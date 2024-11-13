package rest

import (
	"fmt"
	"log"
	"os"

	administrator "github.com/bccfilkom-be/go-server/internal/administrator/interface/rest"
	artikel "github.com/bccfilkom-be/go-server/internal/artikel/interface/rest"
	makanan "github.com/bccfilkom-be/go-server/internal/makanan/interface/rest"
	pengguna "github.com/bccfilkom-be/go-server/internal/pengguna/interface/rest"
	profilPengguna "github.com/bccfilkom-be/go-server/internal/profil_pengguna/interface/rest"
	progresNutrisiHarian "github.com/bccfilkom-be/go-server/internal/progres_nutrisi_harian/interface/rest"
	rekomendasiNutrisiHarian "github.com/bccfilkom-be/go-server/internal/rekomendasi_nutrisi_harian/interface/rest"
	riwayatKesehatan "github.com/bccfilkom-be/go-server/internal/riwayat_kesehatan/interface/rest"

	"github.com/bccfilkom-be/go-server/internal/usecase"
	"github.com/bccfilkom-be/go-server/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/robfig/cron/v3"
)

type Rest struct {
	router                   *gin.Engine
	Usecase                  *usecase.Usecase
	penggunaHandler          *pengguna.PenggunaHandler
	profilPengguna           *profilPengguna.ProfilPenggunaHandler
	riwayatKesehatan         *riwayatKesehatan.RiwayatKesehatanHandler
	rekomendasiNutrisiHarian *rekomendasiNutrisiHarian.RekomendasiNutrisiHarianHandler
	progresNutrisiHarian     *progresNutrisiHarian.ProgresNutrisiHarianHandler
	artikel                  *artikel.ArtikelHandler
	administrator            *administrator.AdministratorHandler
	makanan                  *makanan.MakananHandler
	middleware               middleware.Interface
}

// NewRest constructor untuk Rest, menginisialisasi penggunaHandler juga
func NewRest(usecase *usecase.Usecase, middleware middleware.Interface) *Rest {
	return &Rest{
		router:                   gin.Default(),
		Usecase:                  usecase,
		penggunaHandler:          pengguna.NewPenggunaHandler(usecase),
		profilPengguna:           profilPengguna.NewprofilPenggunaHandler(usecase),
		riwayatKesehatan:         riwayatKesehatan.NewriwayatKesehatanHandler(usecase),
		rekomendasiNutrisiHarian: rekomendasiNutrisiHarian.NewrekomendasiNutrisiHarianHandler(usecase),
		progresNutrisiHarian:     progresNutrisiHarian.NewprogresNutrisiHarianHandler(usecase),
		artikel:                  artikel.NewartikelHandler(usecase),
		administrator:            administrator.NewadministratorHandler(usecase),
		makanan:                  makanan.NewmakananHandler(usecase),
		middleware:               middleware,
	}
}

func (r *Rest) MountEndpoint() {
	r.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))
	r.router.Use(r.middleware.Timeout())

	routerGroup := r.router.Group("/api/v1")

	// Menggunakan handler dari penggunaHandler
	routerGroup.GET("/cek", r.penggunaHandler.GetAllPengguna)
	routerGroup.POST("/register", r.penggunaHandler.PengunaRegister)
	routerGroup.POST("/login", r.penggunaHandler.Login)
	routerGroup.GET("/get-profil", r.middleware.AuthenticateUser, r.profilPengguna.GetProfilPengguna)
	routerGroup.PATCH("/update-profil", r.middleware.AuthenticateUser, r.profilPengguna.UpdateProfilPengguna)
	routerGroup.DELETE("/delete-foto-profil", r.middleware.AuthenticateUser, r.profilPengguna.DeleteFotoProfilPengguna)
	routerGroup.GET("/get-riwayat-kesehatan", r.middleware.AuthenticateUser, r.riwayatKesehatan.GetRiwayatKesehatan)
	routerGroup.PATCH("/update-riwayat-kesehatan", r.middleware.AuthenticateUser, r.riwayatKesehatan.UpdateRiwayatKesehatan)
	routerGroup.DELETE("delete-riwayat-kesehatan", r.middleware.AuthenticateUser, r.riwayatKesehatan.DeleteRiwayatKesehatan)
	routerGroup.PATCH("get-rekomendasi", r.middleware.AuthenticateUser, r.rekomendasiNutrisiHarian.GetRekomendasi)
	routerGroup.POST("create-makanan", r.makanan.CreateMakanan)
	routerGroup.GET("get-makanan", r.makanan.GetMakanan)
	routerGroup.GET("get-progres", r.middleware.AuthenticateUser, r.progresNutrisiHarian.GetProgres)
	routerGroup.PATCH("update-progres", r.middleware.AuthenticateUser, r.progresNutrisiHarian.UpdateProgres)


	// check := cron.New()
	// check.AddFunc("0 0 * * *", func() {
	// 	r.progresNutrisiHarian.ResetProgres()
	// })
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
