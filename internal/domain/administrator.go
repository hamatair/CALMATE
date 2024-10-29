package entity

import (
	"time"
)

type Administrator struct {
	IDAdministrator string    `gorm:"column:id_administrator;primaryKey;type:varchar(255);not null"`
	IDPengguna      string    `gorm:"column:id_pengguna;foreignKey:IDPengguna;references:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;type:varchar(255);not null"`

	NamaAdmin       string    `gorm:"column:nama_admin;type:varchar(255);not null"`

	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
