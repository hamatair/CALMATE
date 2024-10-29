package model

type PenggunaParam struct {
	IDPengguna string
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