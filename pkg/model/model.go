package model

import (
	"mime/multipart"
	"time"
)

type PenggunaParam struct {
	IDPengguna string
	Email      string
	Password   string
}

type PengunaRegister struct {
	Email    string
	Password string

	NamaPengguna      string
	TanggalLahir      time.Time
	JenisKelamin      string
	TinggiBadan       float32
	BeratBadan        float32
	Umur              int
	AktivitasPengguna string
}

type PenggunaLoginResponse struct {
	Token string
}

type ProfilPengguna struct {
	NamaPengguna      string
	TanggalLahir      time.Time
	JenisKelamin      string
	TinggiBadan       float32
	BeratBadan        float32
	Umur              int
	AktivitasPengguna string
	Alamat            string
	NoTeleponPengguna string
	NamaFoto          string
	LinkFoto          string
}

type UpdateRiwayatKesehatan struct {
	Jenis   string
	Detail  string
	Tanggal time.Time
}

type DeleteRiwayatKesehatan struct {
	Jenis string
	Index int
}

type Makanan struct {
	Nama        string
	Jenis       string
	Kalori      float32
	Karbohidrat float32
	Protein     float32
	Lemak       float32
	Gambar      string
}

type GetMakanan struct {
	Nama string
}

type ProgresNutrisiHarian struct {
	JumlahKonsumsiKalori      float32
	JumlahKonsumsiKarbohidrat float32
	JumlahKonsumsiProtein     float32
	JumlahKonsumsiLemak       float32
}

type Foto struct {
	Foto *multipart.FileHeader
}
