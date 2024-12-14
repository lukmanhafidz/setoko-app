package model

import "github.com/google/uuid"

type MProduct struct {
	Id    uuid.UUID
	Name  string
	Qty   int
	Price int
}

func (MProduct) TableName() string {
	return "product.m_product"
}
