package persistence

import (
	"setokoapp/domain/model"
	"setokoapp/domain/repository"

	"gorm.io/gorm"
)

type tTransactionRepository struct {
	db *gorm.DB
}

func NewTTransactionRepository(db *gorm.DB) repository.ITTransaction { //func constructor
	return &tTransactionRepository{db: db}
}

// FindOrderDetails implements repository.ITOrder
func (r *tTransactionRepository) FindOrderReceipt(trxId string) (*model.OrderReceipt, error) {
	orderReceipt := new(model.OrderReceipt)

	err := r.db.Table("transaction.t_transaction tt").
		Select("tt.order_no, "+
			"mm.name as merchant_name, "+
			"mm.link as merchant_link, "+
			"mm.phone as merchant_phone, "+
			"tt.total as total_payment, "+
			"tt.payment_method, "+
			"tt.status as payment_status, "+
			"tt.delivery_method, "+
			"tt.delivery_total").
		Joins("left join merchant.m_merchant mm on tt.merchant_id = mm.id").
		Where("tt.id = ?", trxId).
		Scan(&orderReceipt).Error

	if err != nil {
		return nil, err
	}

	return orderReceipt, nil
}
