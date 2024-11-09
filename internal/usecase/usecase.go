package usecase

import (
	administrator "github.com/bccfilkom-be/go-server/internal/administrator/usecase"
	artikel "github.com/bccfilkom-be/go-server/internal/artikel/usecase"
	pengguna "github.com/bccfilkom-be/go-server/internal/pengguna/usecase"
	profilPengguna "github.com/bccfilkom-be/go-server/internal/profil_pengguna/usecase"
	progresNutrisiHarian "github.com/bccfilkom-be/go-server/internal/progres_nutrisi_harian/usecase"
	rekomendasiNutrisiHarian "github.com/bccfilkom-be/go-server/internal/rekomendasi_nutrisi_harian/usecase"
	riwayatKesehatan "github.com/bccfilkom-be/go-server/internal/riwayat_kesehatan/usecase"

	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
)

type Usecase struct {
	PenggunaUsecase                 pengguna.IPenggunaUsecase
	ProfilPenggunaUsecase           profilPengguna.IProfilPenggunaUsecase
	RiwayatKesehatanUsecase         riwayatKesehatan.IRiwayatKesehatanUsecase
	RekomendasiNutrisiHarianUsecase rekomendasiNutrisiHarian.IRekomendasiNutrisiHarianUsecase
	ProgresNutrisiHarianUsecase     progresNutrisiHarian.IProgresNutrisiHarianUsecase
	ArtikelUsecase                  artikel.IArtikelUsecase
	Administrator                   administrator.IAdministratorUsecase
	// UserService IUserService
	// BookService IBookService
}

type InitParam struct {
	Repository *repository.Repository
	Bcrypt     bcrypt.Interface
	JwtAuth    jwt.Interface
}

func NewUsecase(param InitParam) *Usecase {
	penggunaUsecase := pengguna.NewpenggunaUsecase(*param.Repository, param.Bcrypt, param.JwtAuth)
	profilPenggunaUsecase := profilPengguna.NewProfilPenggunaUsecase(*param.Repository)
	riwayatKesehatan := riwayatKesehatan.NewriwayatKesehatanUsecase(*param.Repository)
	rekomendasiNutrisiHarian := rekomendasiNutrisiHarian.NewrekomendasiNutrisiHarianUsecase(*param.Repository)
	
	return &Usecase{
		PenggunaUsecase: penggunaUsecase,
		ProfilPenggunaUsecase: profilPenggunaUsecase,
		RiwayatKesehatanUsecase: riwayatKesehatan,
		RekomendasiNutrisiHarianUsecase: rekomendasiNutrisiHarian,
	}
}
