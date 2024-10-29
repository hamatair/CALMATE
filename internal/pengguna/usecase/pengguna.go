package usecase

import (
	entity "github.com/bccfilkom-be/go-server/internal/domain"
	"github.com/bccfilkom-be/go-server/internal/pengguna/repository"
	"github.com/bccfilkom-be/go-server/pkg/bcrypt"
	"github.com/bccfilkom-be/go-server/pkg/jwt"
	"github.com/bccfilkom-be/go-server/pkg/model"
)

type IPenggunaUsecase interface {
	// daftar fungsi
	GetPengguna(model.PenggunaParam) (entity.Pengguna, error)
}

type penggunaUsecase struct {
	PenggunaRepository repository.IPenggunaRepository
	bcrypt             bcrypt.Interface
	jwtAuth            jwt.Interface
}

// GetPengguna implements IPenggunaUsecase.
func (*penggunaUsecase) GetPengguna(model.PenggunaParam) (entity.Pengguna, error) {
	panic("unimplemented")
}

func NewpenggunaUsecase(penggunaRepository repository.IPenggunaRepository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) IPenggunaUsecase {
	return &penggunaUsecase{
		PenggunaRepository: penggunaRepository,
		jwtAuth:            jwtAuth,
		bcrypt:             bcrypt,
	}
}
