package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rs/cors"

	"github.com/fguisso/lnstore/api"
)

const port = ":3000"

func main() {
	shutdownChan := make(chan os.Signal)
	signal.Notify(shutdownChan, os.Interrupt)

	r := chi.NewRouter()

	corsMW := cors.Default()
	r.Use(corsMW.Handler)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Logger)

	routes := api.NewRoutes()
	router := api.NewAPIRouter(routes)

	r.Mount("/api", router)

	svr := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		fmt.Println("Listening on port ", port)
		if err := svr.ListenAndServe(); err != nil {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	<-shutdownChan
	fmt.Println("Shutting down the server.")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	svr.Shutdown(ctx)
	defer cancel()
	fmt.Println("Server gracefully stopped.")
}
