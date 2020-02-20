package domain

type Additional struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Bowl struct {
	ID    string `json:"id"`
	Size  int64  `json:"size"`
	Price int64  `json:"price"`
}

type Order struct {
	ID        string `json:"id"`
	Total     int64  `json:"total"`
	Coin      string `json:"coin"`
	LNInvoice string `json:"ln_invoice"`
}

type OrderRequest struct {
	Additionals []string `json:"additionals"`
	Bowl        string   `json:"bowl"`
	Coin        string   `json:"coin"`
}
