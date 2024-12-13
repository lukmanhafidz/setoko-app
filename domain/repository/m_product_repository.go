package repository

import "setokoapp/domain/model"

type IMProduct interface {
	FindProductOrder(orderId string) ([]model.ProductOrder, error)
}
