package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/validators"
)

type TransactionInputDTO struct {
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type TransactionOutputDTO struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userId, err := strconv.Atoi(id)
	if err != nil || userId < 1 || userId > 6 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID is invalid")
		return
	}

	var transaction TransactionInputDTO
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validTransaction, err := validators.NewTransactionValidator(transaction.Valor, transaction.Tipo, transaction.Descricao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := NewTransaction(validTransaction.Valor, validTransaction.Tipo, validTransaction.Descricao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func NewTransaction(valor int64, tipo string, descricao string) (*TransactionOutputDTO, error) {
	return &TransactionOutputDTO{
		Limite: 1000,
		Saldo:  1000,
	}, nil
}
