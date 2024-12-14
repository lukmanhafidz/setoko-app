package model

import "github.com/google/uuid"

type MMerchant struct {
	Id    uuid.UUID
	Name  string
	Link  string
	Phone string
}

func (MMerchant) TableName() string {
	return "merchant.m_merchant"
}
