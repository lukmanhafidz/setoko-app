package usecase

import (
	"errors"
	"setokoapp/constants"
	"setokoapp/domain/mocks"
	"setokoapp/domain/model"
	"setokoapp/utils"
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
	tOrderRepository *mocks.ITOrder
}

func NewGenerateUsecaseTestRepo() *generateUsecaseTestRepo {
	tOrderRepository := new(mocks.ITOrder)

	return &generateUsecaseTestRepo{tOrderRepository: tOrderRepository}
}

func TestGenerateUsecase(t *testing.T) {
	//init config
	constants.MODE_UNIT_TEST = true
	currentTime = utils.GenerateCurrentTime()

	//init usecase with mocked repository
	generateUsecaseTestRepo := NewGenerateUsecaseTestRepo()
	usecase = NewGenerateUsecase(generateUsecaseTestRepo.tOrderRepository)

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
			TotalPrice:   20000,
		},
		{
			ProductName:  "saus",
			Qty:          3,
			ProductPrice: 5000,
			TotalPrice:   15000,
		},
	}

	mockOrderReceipt := model.OrderReceipt{
		MerchantName:     "toko berkah",
		MerchantLink:     "tokoberkah123.com",
		MerchantPhone:    "08123456789",
		TotalPayment:     35000,
		PaymentMethod:    1,
		PaymentStatus:    1,
		OrderNo:          utils.GenTransactionId(),
		OrderAt:          utils.GenerateCurrentTime(),
		DeliveryMethod:   2,
		DeliveryDetail:   "toko berkah | 08123456789",
		SubtotalDelivery: 10000,
	}

	mockResponseData := mockOrderReceipt
	mockResponseData.OrderDetail = mockOrderDetail
	mockResponseData.SubtotalProduct = 35000

	t.Run("Success Generate Receipt", func(t *testing.T) {
		gt.tOrderRepository.On("FindOrderReceipt", orderId).Return(mockOrderReceipt, nil).Once()
		gt.tOrderRepository.On("FindOrderDetails", orderId).Return(mockOrderDetail, nil).Once()

		expectResp := new(model.BaseResp).OK(mockResponseData)
		resp := usecase.GenerateReceipt(orderId.String())
		if assert.NotEqual(t, expectResp, resp) {
			result = false
		}
	})

	return result
}
