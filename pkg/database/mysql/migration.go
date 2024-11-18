package mysql

import (
	"log"

	"github.com/bccfilkom-be/go-server/internal/domain"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	// db.Migrator().DropTable(
	// 	&entity.Pengguna{},
	// 	&entity.ProfilPengguna{},
	// 	&entity.RiwayatKesehatan{},
	// 	&entity.RekomendasiNutrisiHarian{},
	// 	&entity.ProgresNutrisiHarian{},
	// 	&entity.Artikel{},
	// 	&entity.Administrator{},
	// 	&entity.Makanan{},
	// )

	// Migrasi tabel pengguna terlebih dahulu
	if err := db.AutoMigrate(&entity.Pengguna{}); err != nil {
		log.Fatalf("failed migration db for Pengguna: %v", err)
	}

	// Migrasi tabel lainnya setelah pengguna
	if err := db.AutoMigrate(
		&entity.ProfilPengguna{},
		&entity.RiwayatKesehatan{},
		&entity.RekomendasiNutrisiHarian{},
		&entity.ProgresNutrisiHarian{},
		&entity.Artikel{},
		&entity.Administrator{},
		&entity.Makanan{},
	); err != nil {
		log.Fatalf("failed migration db: %v", err)
	}
}
