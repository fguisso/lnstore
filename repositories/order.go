package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type OrderRepository struct {
	db []d.Order
}

var mockOrder = []d.Order{
	d.Order{ID: "rePpNwCxnVkLOJo29W13vTFqzI39", Total: 10000000, Coin: "decred", LNInvoice: "testing"},
}

func (or OrderRepository) NewRepository() *OrderRepository {
	return &OrderRepository{
		db: mockOrder,
	}
}

func (or OrderRepository) ListAll() []d.Order {
	return or.db
}

func (or OrderRepository) FindByID(id string) d.Order {
	for _, item := range or.db {
		if item.ID == id {
			return item
		}
	}
	return d.Order{}
}

func (or *OrderRepository) Create(newOrder d.Order) {
	or.db = append(or.db, newOrder)
}
