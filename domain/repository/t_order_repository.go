package repository

import "setokoapp/domain/model"

type ITOrder interface {
	FindOrderReceipt(orderId string) (*model.OrderReceipt, error)
}
