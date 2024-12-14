package usecases

import (
	"errors"
	"setokoapp/constants"
	"setokoapp/domain/mocks"
	"setokoapp/domain/model"
	"setokoapp/utils"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	currentTime time.Time
	usecase     IGenerateUsecase
	result      = true
	err         = errors.New("error")
	currentDate = utils.GenerateCurrentTime().Format(constants.DATE_TIME_FORMAT)
)

type generateUsecaseTestRepo struct {
	tTransactionRepository *mocks.ITTransaction
	mProductRepository     *mocks.IMProduct
}

func NewGenerateUsecaseTestRepo() *generateUsecaseTestRepo {
	tTransactionRepository := new(mocks.ITTransaction)
	mProductRepository := new(mocks.IMProduct)

	return &generateUsecaseTestRepo{
		tTransactionRepository: tTransactionRepository,
		mProductRepository:     mProductRepository}
}

func TestGenerateUsecase(t *testing.T) {
	//init config
	constants.MODE_UNIT_TEST = true
	currentTime = utils.GenerateCurrentTime()

	//init usecase with mocked repository
	generateUsecaseTestRepo := NewGenerateUsecaseTestRepo()
	usecase = NewGenerateUsecase(
		generateUsecaseTestRepo.tTransactionRepository,
		generateUsecaseTestRepo.mProductRepository)

	//testing
	generateUsecaseTest := generateUsecaseTestRepo.TestGenerate(t)
	assert.Equal(t, true, generateUsecaseTest)
}

func (gt *generateUsecaseTestRepo) TestGenerate(t *testing.T) bool {
	orderId := utils.GenerateNewUUID()

	mockOrderDetail := []model.ProductOrder{ //mocking models
		{
			ProductName:  "kecap",
			Qty:          2,
			ProductPrice: 10000,
			OrderAt:      utils.GenerateCurrentTime(),
		},
		{
			ProductName:  "saus",
			Qty:          3,
			ProductPrice: 5000,
			OrderAt:      utils.GenerateCurrentTime(),
		},
	}

	mockOrderReceipt := model.OrderReceipt{
		OrderNo:          "INV-" + utils.GenTransactionId(),
		MerchantName:     "toko berkah",
		MerchantLink:     "tokoberkah123.com",
		MerchantPhone:    "08123456789",
		TotalPayment:     35000,
		PaymentMethod:    1,
		PaymentStatus:    1,
		DeliveryMethod:   2,
		SubtotalDelivery: 10000,
	}

	mockResponseData := mockOrderReceipt

	subtotalProduct := 0
	var orderAt time.Time
	for _, order := range mockOrderDetail {
		order.TotalPrice = order.Qty * order.ProductPrice
		mockResponseData.OrderDetail = append(mockResponseData.OrderDetail, order)
		subtotalProduct += order.TotalPrice
		orderAt = order.OrderAt
	}

	mockResponseData.DeliveryDetail = strings.Join([]string{mockResponseData.MerchantName, mockResponseData.MerchantPhone}, " | ")
	mockResponseData.SubtotalProduct = subtotalProduct
	mockResponseData.OrderAt = orderAt.Format(constants.DATE_TIME_FORMAT)

	//unit test
	t.Run("Success Generate Receipt", func(t *testing.T) {
		gt.tTransactionRepository.On("FindOrderReceipt", orderId.String()).Return(&mockOrderReceipt, nil).Once() //call mocked repo so it doesnt affect real db
		gt.mProductRepository.On("FindProductOrder", mockOrderReceipt.OrderNo).Return(mockOrderDetail, nil).Once()

		expectResp := new(model.BaseResp).OK(&mockResponseData)
		resp := usecase.GenerateReceipt(orderId.String()) //compare expected response and the actual one
		if !assert.Equal(t, expectResp, resp) {
			result = false
		}
	})

	t.Run("Failed Get Product Order Data", func(t *testing.T) {
		gt.tTransactionRepository.On("FindOrderReceipt", orderId.String()).Return(&mockOrderReceipt, nil).Once()
		gt.mProductRepository.On("FindProductOrder", mockOrderReceipt.OrderNo).Return(nil, err).Once()

		expectResp := new(model.BaseResp).Error(constants.RC_DATA_NOT_FOUND, constants.RD_DATA_NOT_FOUND)
		resp := usecase.GenerateReceipt(orderId.String())
		if !assert.Equal(t, expectResp, resp) {
			result = false
		}
	})

	t.Run("Failed Get Order Receipt Data", func(t *testing.T) {
		gt.tTransactionRepository.On("FindOrderReceipt", orderId.String()).Return(nil, err).Once()

		expectResp := new(model.BaseResp).Error(constants.RC_DATA_NOT_FOUND, constants.RD_DATA_NOT_FOUND)
		resp := usecase.GenerateReceipt(orderId.String())
		if !assert.Equal(t, expectResp, resp) {
			result = false
		}
	})

	return result //return if test success
}
