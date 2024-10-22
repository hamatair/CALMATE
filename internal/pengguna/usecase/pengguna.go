package usecase

type penggunaUsecaseItf interface {}

type penggunaUsecase struct {}

func NewpenggunaUsecase() penggunaUsecaseItf {
    return &penggunaUsecase{}
}
