package entity

import (
	"time"
)

type Makanan struct {
	IDMakanan   string  `gorm:"column:id_makanan;primaryKey"`
	Nama        string  `gorm:"column:nama"`
	Jenis       string  `gorm:"column:jenis"`
	Kalori      float32 `gorm:"column:kalori;type:float"`
	Karbohidrat float32 `gorm:"column:karbohidrat;type:float"`
	Protein     float32 `gorm:"column:protein;type:float"`
	Lemak       float32 `gorm:"column:lemak;type:float"`
	Gambar      string  `gorm:"column:gambar"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
