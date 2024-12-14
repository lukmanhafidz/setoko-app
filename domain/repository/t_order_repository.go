package repository

import "setokoapp/domain/model"

type ITOrder interface {
	FindProductOrder(orderNo string) ([]model.ProductOrder, error)
}
