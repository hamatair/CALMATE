package entity

import "time"

type Artikel struct {
    IDArtikel        string    `gorm:"column:id_artikel;primaryKey;type:varchar(255)"`
    IDPengguna       string    `gorm:"column:id_pengguna;foreignKey:IDPengguna;references:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;type:varchar(255);not null"`

    JudulArtikel     string    `gorm:"column:judul_artikel;type:text"`
    IsiArtikel       string    `gorm:"column:isi_artikel;type:text"`
    TanggalPublikasi time.Time `gorm:"column:tanggal_publikasi"`
    PenulisArtikel   string    `gorm:"column:penulis_artikel;type:varchar(255)"`
    SumberArtikel    string    `gorm:"column:sumber_artikel;type:varchar(255)"`

    CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt        time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
