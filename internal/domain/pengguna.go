package entity

import (
	"time"
)

type Pengguna struct {
	IDPengguna string    `gorm:"column:id_pengguna;primaryKey;type:varchar(255)"`
	Email      string    `gorm:"column:email;type:varchar(255);not null;unique"`
	Password   string    `gorm:"column:password;type:varchar(255);not null"`
	Role       string    `gorm:"column:role;type:varchar(50);not null"`

	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
