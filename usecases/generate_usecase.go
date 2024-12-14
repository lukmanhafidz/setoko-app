package usecases

import (
	"log"
	"setokoapp/constants"
	"setokoapp/domain/model"
	"setokoapp/domain/repository"
	"strings"
	"time"

	"github.com/google/uuid"
)

type generateUsecase struct {
	tTransactionRepository repository.ITTransaction
	mProductRepository     repository.ITOrder
}

type IGenerateUsecase interface {
	GenerateReceipt(trxId string) model.BaseResp
}

func NewGenerateUsecase(
	tTransactionRepository repository.ITTransaction,
	mProductRepository repository.ITOrder) IGenerateUsecase {
	return &generateUsecase{
		tTransactionRepository: tTransactionRepository,
		mProductRepository:     mProductRepository}
}

// GenerateReceipt implements IGenerateUsecase
func (u *generateUsecase) GenerateReceipt(trxId string) model.BaseResp {
	_, err := uuid.Parse(trxId) //validate order Id
	if err != nil {
		log.Println("Uuid Parse Error: ", err)
		return new(model.BaseResp).Error(constants.RC_INVALID_REQUEST_DATA, constants.RD_INVALID_REQUEST_DATA)
	}

	orderReceipt, err := u.tTransactionRepository.FindOrderReceipt(trxId)
	if err != nil || orderReceipt == nil { //if error exist or data is empty
		log.Println("FindOrderReceipt Error: ", err)
		return new(model.BaseResp).Error(constants.RC_DATA_NOT_FOUND, constants.RD_DATA_NOT_FOUND)
	}

	orderDetail, err := u.mProductRepository.FindProductOrder(orderReceipt.OrderNo)
	if err != nil || len(orderDetail) <= 0 { //if error exist or data is empty
		log.Println("FindProductOrder Error: ", err)
		return new(model.BaseResp).Error(constants.RC_DATA_NOT_FOUND, constants.RD_DATA_NOT_FOUND)
	}

	subtotalProduct := 0
	var orderAt time.Time
	for _, order := range orderDetail {
		order.TotalPrice = order.Qty * order.ProductPrice
		orderReceipt.OrderDetail = append(orderReceipt.OrderDetail, order)
		subtotalProduct += order.TotalPrice
		orderAt = order.OrderAt
	}

	orderReceipt.DeliveryDetail = strings.Join([]string{orderReceipt.MerchantName, orderReceipt.MerchantPhone}, " | ")
	orderReceipt.OrderAt = orderAt.Format(constants.DATE_TIME_FORMAT)
	orderReceipt.SubtotalProduct = subtotalProduct

	return new(model.BaseResp).OK(orderReceipt)
}
