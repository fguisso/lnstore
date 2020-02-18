package api

import (
	"fmt"
	"net/http"

	r "github.com/fguisso/lnstore/repositories"
)

type Routes struct {
	ar *r.AdditionalRepository
	br *r.BowlRepository
	or *r.OrderRepository
}

func NewRoutes() *Routes {
	return &Routes{
		ar: r.Mocked.Additionals,
		br: r.Mocked.Bowls,
		or: r.Mocked.Orders,
	}
}

func (routes *Routes) GetAdditionals(w http.ResponseWriter, _ *http.Request) {
	additionalList := routes.ar.ListAll()

	if len(additionalList) == 0 {
		w.Write([]byte("Additional list empty!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", additionalList)))
}

func (routes *Routes) GetBowls(w http.ResponseWriter, _ *http.Request) {
	bowlList := routes.br.ListAll()

	if len(bowlList) == 0 {
		w.Write([]byte("Bowl list empty!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", bowlList)))
}

func (routes *Routes) GetOrders(w http.ResponseWriter, _ *http.Request) {
	orderList := routes.or.ListAll()

	if len(orderList) == 0 {
		w.Write([]byte("Order list empty!"))
		return
	}

	w.Write([]byte(fmt.Sprintf("%v", orderList)))
}
