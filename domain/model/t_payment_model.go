package model

import "github.com/google/uuid"

type TPayment struct {
	Id     uuid.UUID
	Method int
	Status int
	Total  int
}

func (TPayment) TableName() string {
	return "payment.t_payment"
}
