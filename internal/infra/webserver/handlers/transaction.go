package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	usecase "github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/domain/use-case"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/validators"
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	dbValue := r.Context().Value("DB")
	if dbValue == nil {
		http.Error(w, "Conexão do banco de dados não encontrada no contexto", http.StatusInternalServerError)
		return
	}

	db, ok := dbValue.(*sql.DB)
	if !ok {
		http.Error(w, "Valor no contexto não pode ser convertido para *sql.DB", http.StatusInternalServerError)
		return
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil || userId < 1 || userId > 5 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User ID is invalid")
		return
	}

	var transaction dto.TransactionInputDTO
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validTransaction, err := validators.TransactionValidator(transaction.Valor, transaction.Tipo, transaction.Descricao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := usecase.NewTransactionUseCase(
		db,
		validTransaction.Valor,
		validTransaction.Tipo,
		validTransaction.Descricao,
		userId,
	)
	if err != nil {
		if err.Error() == "saldo insuficiente" {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
