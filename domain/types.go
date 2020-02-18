package domain

type Additional struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Bowl struct {
	ID    int32   `json:"id"`
	Size  int64   `json:"size"`
	Price float64 `json:"price"`
}

type Order struct {
	ID        int32   `json:"id"`
	Total     float64 `json:"total"`
	Coin      string  `json:"coin"`
	LNInvoice string  `json:"ln_invoice"`
}
