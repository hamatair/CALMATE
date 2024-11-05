package model

import "time"

type PenggunaParam struct {
	IDPengguna string
	Email      string
	Password   string
}

type PengunaRegister struct {
	Email    string
	Password string

	NamaPengguna      string
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
	Foto              string
}

type UpdateRiwayatKesehatan struct {
	Jenis string
	Detail string
	Tanggal time.Time
}

type DeleteRiwayatKesehatan struct {
	Jenis string
	Index int
}

