package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type AdditionalRepository struct {
	db []d.Additional
}

var mock = []d.Additional{
	d.Additional{ID: 1, Name: "Banana", Price: 6000},
	d.Additional{ID: 2, Name: "Guarana", Price: 4000},
	d.Additional{ID: 3, Name: "Granola", Price: 5000},
	d.Additional{ID: 4, Name: "Honey", Price: 100},
}

func (ar AdditionalRepository) NewRepository() *AdditionalRepository {
	return &AdditionalRepository{
		db: mock,
	}
}

func (ar AdditionalRepository) ListAll() []d.Additional {
	return ar.db
}

func (ar AdditionalRepository) FindByID(id int32) d.Additional {
	for _, item := range ar.db {
		if item.ID == id {
			return item
		}
	}
	return d.Additional{}
}
