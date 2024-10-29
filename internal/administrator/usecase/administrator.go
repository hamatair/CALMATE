package usecase

import "github.com/bccfilkom-be/go-server/internal/administrator/repository"

type IAdministratorUsecase interface {}

type administratorUsecase struct {
    AdministratorRepository repository.IAdministratorRepository
}

func NewadministratorUsecase(administratorRepository repository.IAdministratorRepository) IAdministratorUsecase {
    return &administratorUsecase{
        AdministratorRepository: administratorRepository,
    }
}
