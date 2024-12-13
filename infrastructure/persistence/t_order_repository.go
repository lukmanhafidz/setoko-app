package persistence

import (
	"setokoapp/domain/model"
	"setokoapp/domain/repository"

	"gorm.io/gorm"
)

type tOrderRepository struct {
	db *gorm.DB
}

func NewTOrderRepository(db *gorm.DB) repository.ITOrder { //func constructor
	return &tOrderRepository{db: db}
}

// FindOrderDetails implements repository.ITOrder
func (r *tOrderRepository) FindOrderDetails(orderId string) (model.OrderReceipt, error) {
	r.db.Model(&model.TOrder{}).
		Select("").
		Joins("")
}
