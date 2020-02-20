package repositories

import (
	d "github.com/fguisso/lnstore/domain"
)

type BowlRepository struct {
	db []d.Bowl
}

var mockBowls = []d.Bowl{
	d.Bowl{ID: "M5sDcTYjR2q9TmCVvtolDxokUulI", Size: 150, Price: 2000},
	d.Bowl{ID: "dCLgji8toHHVl90NoQaGVmis19ce", Size: 250, Price: 2500},
	d.Bowl{ID: "5w66UAfEGLtbeacqilba53H5NVLr", Size: 500, Price: 5000},
	d.Bowl{ID: "KEj0XjGaNqjhRsqgf93jV1P4qi7r", Size: 1000, Price: 8000},
}

func (bd BowlRepository) NewRepository() *BowlRepository {
	return &BowlRepository{
		db: mockBowls,
	}
}

func (bd BowlRepository) ListAll() []d.Bowl {
	return bd.db
}

func (bd BowlRepository) FindByID(id string) d.Bowl {
	for _, item := range bd.db {
		if item.ID == id {
			return item
		}
	}
	return d.Bowl{}
}
