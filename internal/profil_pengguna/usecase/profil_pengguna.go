package usecase

import(
    "github.com/bccfilkom-be/go-server/internal/profil_pengguna/repository"
)

type IProfilPenggunaUsecase interface {}

type profilPenggunaUsecase struct {
    ProfilPenggunaRepository repository.IProfilPenggunaRepository
}

func NewProfilPenggunaUsecase(profilPenggunaRepository repository.IProfilPenggunaRepository) IProfilPenggunaUsecase {
    return &profilPenggunaUsecase{
        ProfilPenggunaRepository: profilPenggunaRepository,
    }
}
