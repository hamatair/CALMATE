package usecase

import "github.com/bccfilkom-be/go-server/internal/riwayat_kesehatan/repository"

type IRiwayatKesehatanUsecase interface {}

type riwayatKesehatanUsecase struct {
    RiwayatKesehatanRepository repository.IRiwayatKesehatanRepository
}

func NewriwayatKesehatanUsecase(riwayatKesehatanRepository repository.IRiwayatKesehatanRepository) IRiwayatKesehatanUsecase {
    return &riwayatKesehatanUsecase{
        RiwayatKesehatanRepository: riwayatKesehatanRepository,
    }
}
