package usecase

import (
	"database/sql"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

func NewExtractUseCase(db *sql.DB, userId int64) (*dto.ExtractOutputDTO, error) {
	transactions, err := database.GetLastTransactionsByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	balance, err := database.GetBalanceAndLimitByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	return &dto.ExtractOutputDTO{
		Balance:          *balance,
		LastTransactions: transactions,
	}, nil
}
