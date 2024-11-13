package entity

import "time"

type ProgresNutrisiHarian struct {
	IDProgresNutrisiHarian    string  `gorm:"column:id_progres_nutrisi_harian;primaryKey;type:varchar(255)"`
	IDPengguna                string  `gorm:"column:id_pengguna;foreignKey:IDPengguna;references:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;type:varchar(255);not null"`
	JumlahKonsumsiKalori      float32 `gorm:"column:jumlah_konsumsi_kalori;type:float"`
	JumlahKonsumsiKarbohidrat float32 `gorm:"column:jumlah_konsumsi_karbohidrat;type:float"`
	JumlahKonsumsiProtein     float32 `gorm:"column:jumlah_konsumsi_protein;type:float"`
	JumlahKonsumsiLemak       float32 `gorm:"column:jumlah_konsumsi_lemak;type:float"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
