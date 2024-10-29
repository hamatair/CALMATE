package entity

import "time"

type RekomendasiNutrisiHarian struct {
	IDRekomendasi           string  `gorm:"column:id_rekomendasi;primaryKey;type:varchar(255)"`
	IDPengguna              string  `gorm:"column:id_pengguna;foreignKey:IDPengguna;references:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;type:varchar(255);not null"`
	JumlahKaloriHarian      float32 `gorm:"column:jumlah_kalori_harian;type:float"`
	AsupanKarbohidratHarian float32 `gorm:"column:asupan_karbohidrat_harian;type:float"`
	AsupanProteinHarian     float32 `gorm:"column:asupan_protein_harian;type:float"`
	AsupanLemakHarian       float32 `gorm:"column:asupan_lemak_harian;type:float"`
	AsupanVitaminHarian     float32 `gorm:"column:asupan_vitamin_harian;type:float"`
	AsupanMineralHarian     float32 `gorm:"column:asupan_mineral_harian;type:float"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
