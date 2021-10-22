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
	mux.Use(LoadSession)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.ShowHomePage)
	mux.Get("/customers", handlers.Repo.ShowAllCustomers)
	mux.Post("/customers/search", handlers.Repo.SearchCustomers)

	mux.Get("/customer/{id}/view", handlers.Repo.ShowCustomer)
	mux.Get("/customer/{id}", handlers.Repo.ShowCustomerForm)
	mux.Post("/customer/{id}", handlers.Repo.EditCustomer)

	mux.Get("/customer", handlers.Repo.ShowCustomerForm)
	mux.Post("/customer", handlers.Repo.AddCustomer)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
