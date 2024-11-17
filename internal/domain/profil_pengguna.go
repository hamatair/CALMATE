package entity

import "time"

type ProfilPengguna struct {
	IDProfil   string `gorm:"column:id_profil;primaryKey;type:varchar(255)"`
	IDPengguna string `gorm:"column:id_pengguna;type:varchar(255);not null"`

	NamaPengguna      string    `gorm:"column:nama_pengguna;type:varchar(255);not null"`
	TanggalLahir      time.Time `gorm:"column:tanggal_lahir;type:date"`
	JenisKelamin      string    `gorm:"column:jenis_kelamin;type:varchar(1)"`
	TinggiBadan       float32   `gorm:"column:tinggi_badan;type:float"`
	BeratBadan        float32   `gorm:"column:berat_badan;type:float"`
	Umur              int       `gorm:"column:umur;type:int"`
	AktivitasPengguna string    `gorm:"column:aktivitas_pengguna;type:varchar(255)"`
	Alamat            string    `gorm:"column:alamat;type:text"`
	NoTeleponPengguna string    `gorm:"column:no_telepon_pengguna;type:varchar(20)"`
	NamaFoto          string    `gorm:"column:nama_foto;type:varchar(255)"`
	LinkFoto          string    `gorm:"column:link_foto;type:varchar(255)"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}
