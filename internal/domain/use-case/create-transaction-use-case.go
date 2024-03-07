package usecase

import (
	"database/sql"
	"errors"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/domain/entity"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

func NewTransactionUseCase(db *sql.DB, valor int64, tipo string, descricao string, userId int64) (*dto.TransactionOutputDTO, error) {
	entityTransaction := entity.NewTransaction(valor, tipo, descricao)

	balance, err := database.GetBalanceAndLimitByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	if balance.Total+balance.Limit <= entityTransaction.Valor && entityTransaction.Tipo == "d" {
		return nil, errors.New("saldo insuficiente")
	}

	go func() {
		_ = database.CreateTransaction(db, &dto.TransactionInputDTO{
			Valor:     entityTransaction.Valor,
			Tipo:      entityTransaction.Tipo,
			Descricao: entityTransaction.Descricao,
			ClienteID: userId,
		})
	}()

	if tipo == "c" {
		balance.Total = balance.Total + entityTransaction.Valor
	} else {
		balance.Total = balance.Total - entityTransaction.Valor
	}

	return &dto.TransactionOutputDTO{
		Limite: balance.Limit,
		Saldo:  balance.Total,
	}, nil
}
