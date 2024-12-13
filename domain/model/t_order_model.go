package model

import (
	"time"

	"github.com/google/uuid"
)

type TOrder struct {
	Id             uuid.UUID
	PaymentId      uuid.UUID
	MerchantId     uuid.UUID
	ProductId      uuid.UUID
	OrderNo        string
	OrderAt        time.Time
	DeliveryMethod int //1.makan di tempat 2.ambil sendiri 3.diantar kurir
}

func (TOrder) TableName() string {
	return "order.t_order"
}
