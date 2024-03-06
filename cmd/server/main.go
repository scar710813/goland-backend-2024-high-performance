package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/webserver/handlers"
)

func main() {
	db, err := database.NewMySQLStorage()
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(20000)
	db.SetMaxIdleConns(10000)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "DB", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	r.Get("/clientes/{id}/extrato", handlers.ExtractHandler)
	r.Post("/clientes/{id}/transacoes", handlers.TransactionHandler)

	http.ListenAndServe(":8080", r)
}
