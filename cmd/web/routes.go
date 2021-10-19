package main

import (
	"github.com/ekateryna-tln/wallester_task/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.ShowAllCustomers)
	mux.Get("/customers/{id}", handlers.Repo.ShowCustomer)
	mux.Get("/customers", handlers.Repo.AddCustomer)
	mux.Post("/customers/{id}", handlers.Repo.EditCustomer)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
