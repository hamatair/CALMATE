package repository

import (
	administrator "github.com/bccfilkom-be/go-server/internal/administrator/repository"
	artikel "github.com/bccfilkom-be/go-server/internal/artikel/repository"
	makanan "github.com/bccfilkom-be/go-server/internal/makanan/repository"
	pengguna "github.com/bccfilkom-be/go-server/internal/pengguna/repository"
	profilPengguna "github.com/bccfilkom-be/go-server/internal/profil_pengguna/repository"
	progresNutrisiHarian "github.com/bccfilkom-be/go-server/internal/progres_nutrisi_harian/repository"
	rekomendasiNutrisiHarian "github.com/bccfilkom-be/go-server/internal/rekomendasi_nutrisi_harian/repository"
	riwayatKesehatan "github.com/bccfilkom-be/go-server/internal/riwayat_kesehatan/repository"

	"gorm.io/gorm"
)

type Repository struct {
	PenggunaRepository                 pengguna.IPenggunaRepository
	ProfilPenggunaRepository           profilPengguna.IProfilPenggunaRepository
	RiwayatKesehatanRepository         riwayatKesehatan.IRiwayatKesehatanRepository
	RekomendasiNutrisiHarianRepository rekomendasiNutrisiHarian.IRekomendasiNutrisiharianRepository
	ProgresNutrisiHarian               progresNutrisiHarian.IProgresNutrisiHarianRepository
	Artikel                            artikel.IArtikelRepository
	Administrator                      administrator.IAdministratorRepository
	Makanan                            makanan.IMakananRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		PenggunaRepository:                 pengguna.NewPenggunaRepository(db),
		ProfilPenggunaRepository:           profilPengguna.NewProfilPenggunaRepository(db),
		RiwayatKesehatanRepository:         riwayatKesehatan.NewRiwayatKesehatanRepository(db),
		RekomendasiNutrisiHarianRepository: rekomendasiNutrisiHarian.NewRekomendasiNutrisiHarianRepository(db),
		ProgresNutrisiHarian:               progresNutrisiHarian.NewProgresNutrisiHarianRepository(db),
		Artikel:                            artikel.NewArtikelRepository(db),
		Administrator:                      administrator.NewAdministratorRepository(db),
		Makanan:                            makanan.NewMakananRepository(db),
	}
}
