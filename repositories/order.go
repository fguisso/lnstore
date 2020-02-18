package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type OrderRepository struct {
	db []d.Order
}

var mockOrder = []d.Order{
	d.Order{ID: 10, Total: 10000000, Coin: "decred", LNInvoice: "testing"},
}

func (or OrderRepository) NewRepository() *OrderRepository {
	return &OrderRepository{
		db: mockOrder,
	}
}

func (or OrderRepository) ListAll() []d.Order {
	return or.db
}

func (or OrderRepository) FindByID(id int32) d.Order {
	for _, item := range or.db {
		if item.ID == id {
			return item
		}
	}
	return d.Order{}
}
