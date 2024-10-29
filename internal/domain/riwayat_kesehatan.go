package entity

import (
    "time"
    
)

type RiwayatKesehatan struct {
    IDRiwayat       string         `gorm:"column:id_riwayat;primaryKey;type:varchar(36)"`
    IDPengguna      string         `gorm:"column:id_pengguna;foreignKey:IDPengguna;references:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;type:varchar(36);not null"`

    NilaiBMI        float32        `gorm:"column:nilai_bmi;type:float"`
    Alergi          string         `gorm:"column:alergi;type:text"`             // Menyimpan sebagai JSON atau TEXT
    RiwayatObat     string         `gorm:"column:riwayat_obat;type:text"`       // Menyimpan sebagai JSON atau TEXT
    RiwayatOperasi  string         `gorm:"column:riwayat_operasi;type:text"`    // Menyimpan sebagai JSON atau TEXT
    RiwayatPenyakit string         `gorm:"column:riwayat_penyakit;type:text"`   // Menyimpan sebagai JSON atau TEXT

    CreatedAt       time.Time      `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt       time.Time      `gorm:"column:updated_at;autoUpdateTime"`
}
