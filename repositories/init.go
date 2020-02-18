package repositories

type NewMock struct {
	Additionals *AdditionalRepository
	Bowls       *BowlRepository
	Orders      *OrderRepository
}

var Mocked = NewMock{
	Additionals: AdditionalRepository{}.NewRepository(),
	Bowls:       BowlRepository{}.NewRepository(),
	Orders:      OrderRepository{}.NewRepository(),
}
