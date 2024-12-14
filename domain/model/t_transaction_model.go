package model

import "github.com/google/uuid"

type TPayment struct {
	Id             uuid.UUID
	Method         int
	Status         int
	Total          int
	MerchantId     uuid.UUID
	OrderNo        string
	DeliveryMethod int //1.makan di tempat 2.ambil sendiri 3.diantar kurir
	DeliveryTotal  int
}

func (TPayment) TableName() string {
	return "transaction.t_transaction"
}
