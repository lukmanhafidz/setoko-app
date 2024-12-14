package persistence

import (
	"setokoapp/domain/model"
	"setokoapp/domain/repository"

	"gorm.io/gorm"
)

type tOrderRepository struct {
	db *gorm.DB
}

func NewTOrderRepository(db *gorm.DB) repository.ITOrder {
	return &tOrderRepository{db: db}
}

// FindOrderDetail implements repository.IMProduct
func (r *tOrderRepository) FindProductOrder(orderNo string) ([]model.ProductOrder, error) {
	productOrders := []model.ProductOrder{}

	err := r.db.Table(`"order".t_order tor`).
		Select("tor.order_no, "+
			"tor.order_at, "+
			"tor.qty, "+
			"mp.name as product_name, "+
			"mp.price as product_price").
		Where("tor.order_no = ?", orderNo).
		Joins("left join product.m_product mp on tor.product_id = mp.id").
		Scan(&productOrders).Error

	if err != nil {
		return nil, err
	}

	return productOrders, nil
}
