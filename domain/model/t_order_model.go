package model

import (
	"time"

	"github.com/google/uuid"
)

type TOrder struct {
	Id        uuid.UUID
	ProductId uuid.UUID
	OrderNo   string
	OrderAt   time.Time
}

func (TOrder) TableName() string {
	return "order.t_order"
}
