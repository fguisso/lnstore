package repositories

type NewMock struct {
	Additionals *AdditionalRepository
	Bowls       *BowlRepository
}

var Mocked = NewMock{
	Additionals: AdditionalRepository{}.NewRepository(),
	Bowls:       BowlRepository{}.NewRepository(),
}
