package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	"github.com/fguisso/lnstore/api"
)

func main() {
	r := chi.NewRouter()

	corsMW := cors.Default()
	r.Use(corsMW.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	routes := api.NewRoutes()
	router := api.NewAPIRouter(routes)

	r.Mount("/api", router)

	http.ListenAndServe(":3000", r)
}
