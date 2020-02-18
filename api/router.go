package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func NewAPIRouter(routes *Routes) chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "lnstore api running")
	})

	r.Route("/additionals", func(r chi.Router) {
		r.Get("/", routes.GetAdditionals)
	})

	return r
}