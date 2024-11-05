package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type RiwayatKesehatan struct {
	IDRiwayat  string `gorm:"column:id_riwayat;primaryKey;type:varchar(36)"`
	IDPengguna string `gorm:"column:id_pengguna;type:varchar(36);not null"`

	NilaiBMI        float32 `gorm:"column:nilai_bmi;type:float"`
	Alergi          Detail  `gorm:"column:alergi;type:json"`
	RiwayatObat     Detail  `gorm:"column:riwayat_obat;type:json"`
	RiwayatOperasi  Detail  `gorm:"column:riwayat_operasi;type:json"`
	RiwayatPenyakit Detail  `gorm:"column:riwayat_penyakit;type:json"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type DetailRiwayatKesehatan struct {
	ID      string
	Detail  string
	Tanggal time.Time
}

func (r DetailRiwayatKesehatan) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *DetailRiwayatKesehatan) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &r)
}

type Detail []DetailRiwayatKesehatan

func (a Detail) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Detail) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}
