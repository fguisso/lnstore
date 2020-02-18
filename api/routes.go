package api

import (
	"encoding/json"
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
func writeJSON(w http.ResponseWriter, thing interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(thing); err != nil {
		fmt.Println("deu ruim")
	}
}

func (routes *Routes) GetAdditionals(w http.ResponseWriter, _ *http.Request) {
	additionalList := routes.ar.ListAll()

	if len(additionalList) == 0 {
		w.Write([]byte("Additional list empty!"))
		return
	}

	writeJSON(w, additionalList)
}

func (routes *Routes) GetBowls(w http.ResponseWriter, _ *http.Request) {
	bowlList := routes.br.ListAll()

	if len(bowlList) == 0 {
		w.Write([]byte("Bowl list empty!"))
		return
	}

	writeJSON(w, bowlList)
}

func (routes *Routes) GetOrders(w http.ResponseWriter, _ *http.Request) {
	orderList := routes.or.ListAll()

	if len(orderList) == 0 {
		w.Write([]byte("Order list empty!"))
		return
	}

	writeJSON(w, orderList)
}
