package usecase

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/repository"
	"github.com/bccfilkom-be/go-server/pkg/model"
	"github.com/google/uuid"
)

type IMakananUsecase interface {
	CreateMakanan(model.Makanan) error
	GetMakanan(string) ([]entity.Makanan, error)
}

type makananUsecase struct {
	Repository repository.Repository
}

// GetMakanan implements IMakananUsecase.
func (u *makananUsecase) GetMakanan(jenis string) ([]entity.Makanan, error) {
	makanan, err := u.Repository.Makanan.GetMakanan(jenis)
	if err != nil {
		return []entity.Makanan{}, err
	}

	return makanan, err
}

// CreateMakanan implements IMakananUsecase.
func (u *makananUsecase) CreateMakanan(param model.Makanan) error {
	makanan := entity.Makanan{
		IDMakanan:   uuid.New().String(),
		Jenis:       param.Jenis,
		Nama:        param.Nama,
		Kalori:      param.Kalori,
		Karbohidrat: param.Karbohidrat,
		Protein:     param.Protein,
		Lemak:       param.Lemak,
		Gambar:      param.Gambar,
	}

	err := u.Repository.Makanan.CreateMakanan(makanan)
	if err != nil {
		return err
	}

	return err

}

func NewmakananUsecase(repository repository.Repository) IMakananUsecase {
	return &makananUsecase{
		Repository: repository,
	}
}
