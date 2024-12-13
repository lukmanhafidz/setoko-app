package repository

import "setokoapp/domain/model"

type ITOrder interface {
	FindOrderDetails(orderId string) (model.OrderReceipt, error)
}
