package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"

	d "github.com/fguisso/lnstore/domain"
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
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(thing); err != nil {
		fmt.Println("deu ruim")
	}
}

func randomID() string {
	var safeCharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	rand.Seed(time.Now().UnixNano())
	id := make([]byte, 28)
	for i := range id {
		id[i] = safeCharset[rand.Intn(len(safeCharset))]
	}
	return string(id)
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

func (routes *Routes) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "orderID"))
	orderItem := routes.or.FindByID(string(id))
	if orderItem.ID == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Order not found!"))
		return
	}

	writeJSON(w, orderItem)
}

func (routes *Routes) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrderRequest d.OrderRequest
	var newOrder d.Order

	err := json.NewDecoder(r.Body).Decode(&newOrderRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newOrder.ID = randomID()
	newOrder.Coin = newOrderRequest.Coin
	newOrder.LNInvoice = "dcrln10000msdfjlkawhuaskhASDFasdf"

	for _, itemID := range newOrderRequest.Additionals {
		additional := routes.ar.FindByID(itemID)
		newOrder.Total += additional.Price
	}

	bowl := routes.br.FindByID(newOrderRequest.Bowl)
	newOrder.Total += bowl.Price

	routes.or.Create(newOrder)
	writeJSON(w, newOrder)
}
