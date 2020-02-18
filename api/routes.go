package api

import (
	"fmt"
	"net/http"

	r "github.com/fguisso/lnstore/repositories"
)

type Routes struct {
	as *r.AdditionalRepository
}

func NewRoutes(as *r.AdditionalRepository) *Routes {
	return &Routes{
		as: as,
	}
}

func (routes *Routes) GetAdditionals(w http.ResponseWriter, _ *http.Request) {
	additionalList := routes.as.ListAll()

	if len(additionalList) == 0 {
		w.Write([]byte("Additional list empty!"))
	}

	w.Write([]byte(fmt.Sprintf("%v", additionalList)))
}
