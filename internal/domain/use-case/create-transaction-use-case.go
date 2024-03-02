package usecase

import (
	"database/sql"
	"errors"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

func NewTransactionUseCase(db *sql.DB, valor int64, tipo string, descricao string, userId int64) (*dto.TransactionOutputDTO, error) {
	balance, err := database.GetBalanceAndLimitByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	if tipo == "d" && valor >= (balance.Total+balance.Limit) {
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

	var saldo int64
	if tipo == "c" {
		saldo = balance.Total + valor
	} else {
		saldo = balance.Total - valor
	}

	return &dto.TransactionOutputDTO{
		Limite: balance.Limit,
		Saldo:  saldo,
	}, nil
}
