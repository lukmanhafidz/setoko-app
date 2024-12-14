package repository

import "setokoapp/domain/model"

type ITTransaction interface {
	FindOrderReceipt(trxId string) (*model.OrderReceipt, error)
}
