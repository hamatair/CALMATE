package usecase

import "github.com/bccfilkom-be/go-server/internal/artikel/repository"

type IArtikelUsecase interface {}

type artikelUsecase struct {
    ArtikelRepository repository.IArtikelRepository
}

func NewartikelUsecase(artikelRepository repository.IArtikelRepository) IArtikelUsecase {
    return &artikelUsecase{
        ArtikelRepository: artikelRepository,
    }
}
