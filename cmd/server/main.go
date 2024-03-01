package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/webserver/handlers"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/clientes/{id}/extrato", handlers.ExtractHandler)
	r.Post("/clientes/{id}/transacoes", handlers.TransactionHandler)

	http.ListenAndServe(":8080", r)
}
