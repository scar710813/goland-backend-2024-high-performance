package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("User ID is invalid")
		return
	}

	var transaction dto.TransactionInputDTO
	err = json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validTransaction, err := TransactionValidator(transaction.Valor, transaction.Tipo, transaction.Descricao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := NewTransactionUseCase(
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

func TransactionValidator(value int64, tipo string, description string) (*dto.TransactionInputDTO, error) {
	if value <= 0 {
		return nil, errors.New("valor da transação precisa ser positivo")
	}
	if tipo != "c" && tipo != "d" {
		return nil, errors.New("tipo da transação inválido")
	}
	if len(description) < 1 || len(description) > 10 {
		return nil, errors.New("descrição da transação deve ter entre 1 e 10 caracteres")
	}

	return &dto.TransactionInputDTO{
		Valor:     value,
		Tipo:      tipo,
		Descricao: description,
	}, nil
}

func NewTransactionUseCase(db *sql.DB, valor int64, tipo string, descricao string, userId int64) (*dto.TransactionOutputDTO, error) {
	balance, err := database.GetBalanceAndLimitByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	if tipo == "d" && valor > (balance.Total+balance.Limit) {
		return nil, errors.New("saldo insuficiente")
	}

	err = database.CreateTransaction(db, &dto.TransactionInputDTO{
		Valor:     valor,
		Tipo:      tipo,
		Descricao: descricao,
		ClienteID: userId,
	})
	if err != nil {
		return nil, err
	}

	return &dto.TransactionOutputDTO{
		Limite: balance.Limit,
		Saldo:  balance.Total,
	}, nil
}
