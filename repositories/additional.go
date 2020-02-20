package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type AdditionalRepository struct {
	db []d.Additional
}

var mock = []d.Additional{
	d.Additional{ID: "LOY0x8anmjskGlB5MiinOSRUUdD3", Name: "Banana", Price: 6000},
	d.Additional{ID: "pvlbNsEJOluIZeW0qzvsvE4PYSME", Name: "Guarana", Price: 4000},
	d.Additional{ID: "nUPEDfV2Kr8DxkoB9dknqj5Km1Da", Name: "Granola", Price: 5000},
	d.Additional{ID: "Q69VVVzwLPvrGP4cSelJAz4iR87n", Name: "Honey", Price: 100},
}

func (ar AdditionalRepository) NewRepository() *AdditionalRepository {
	return &AdditionalRepository{
		db: mock,
	}
}

func (ar AdditionalRepository) ListAll() []d.Additional {
	return ar.db
}

func (ar AdditionalRepository) FindByID(id string) d.Additional {
	for _, item := range ar.db {
		if item.ID == id {
			return item
		}
	}
	return d.Additional{}
}
