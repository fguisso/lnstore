package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type BowlRepository struct {
	db []d.Bowl
}

var mockBowls = []d.Bowl{
	d.Bowl{ID: 1, Size: 150, Price: 2000},
	d.Bowl{ID: 2, Size: 250, Price: 2500},
	d.Bowl{ID: 3, Size: 500, Price: 5000},
	d.Bowl{ID: 4, Size: 1000, Price: 8000},
}

func (bd BowlRepository) NewRepository() *BowlRepository {
	return &BowlRepository{
		db: mockBowls,
	}
}

func (bd BowlRepository) ListAll() []d.Bowl {
	return bd.db
}

func (bd BowlRepository) FindByID(id int32) d.Bowl {
	for _, item := range bd.db {
		if item.ID == id {
			return item
		}
	}
	return d.Bowl{}
}
