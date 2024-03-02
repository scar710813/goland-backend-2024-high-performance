package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

func ExtractHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil || userId < 1 || userId > 5 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID is invalid")
		return
	}

	result, err := NewExtractUseCase(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func NewExtractUseCase(userId int64) (*dto.ExtractOutputDTO, error) {
	db, err := database.NewMySQLStorage()
	if err != nil {
		return nil, err
	}

	transactions, err := database.GetLastTransactionsByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	balance, err := database.GetBalance(db, userId)
	if err != nil {
		return nil, err
	}

	limit, err := database.GetLimitByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	saldo := dto.Balance{
		Total:     balance,
		CreatedAt: time.Now().Format("2006-01-02T15:04:05.999999Z"),
		Limit:     limit,
	}

	return &dto.ExtractOutputDTO{
		Balance:          saldo,
		LastTransactions: transactions,
	}, nil
}
