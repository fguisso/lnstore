package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	"github.com/fguisso/lnstore/api"
	repo "github.com/fguisso/lnstore/repositories"
)

func main() {
	r := chi.NewRouter()

	corsMW := cors.Default()
	r.Use(corsMW.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/", func(r chi.Router) {
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("[{'name': 'granola'}]"))
		})
	})

	as := &repo.AdditionalRepository{}
	routes := api.NewRoutes(as.NewRepository())
	router := api.NewAPIRouter(routes)

	r.Mount("/api", router)

	http.ListenAndServe(":3000", r)
}
