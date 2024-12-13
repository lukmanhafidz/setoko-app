package usecase

import (
	"setokoapp/domain/model"
	"setokoapp/domain/repository"
)

type generateUsecase struct {
	tOrderRepository repository.ITOrder
}

type IGenerateUsecase interface {
	GenerateReceipt(orderId string) model.BaseResp
}

func NewGenerateUsecase(tOrderRepository repository.ITOrder) IGenerateUsecase {
	return &generateUsecase{tOrderRepository: tOrderRepository}
}

// GenerateReceipt implements IGenerateUsecase
func (u *generateUsecase) GenerateReceipt(orderId string) model.BaseResp {
	panic("unimplemented")
}
