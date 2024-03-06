package usecase

import (
	"database/sql"
	"sync"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/database"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

func NewExtractUseCase(db *sql.DB, userId int64) (*dto.ExtractOutputDTO, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var resultErr error

	var transactions *[]dto.LastTransaction
	var balance *dto.Balance

	wg.Add(2)

	go func() {
		defer wg.Done()
		t, err := database.GetLastTransactionsByUserId(db, userId)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			resultErr = err
			return
		}
		transactions = &t
	}()

	go func() {
		defer wg.Done()
		b, err := database.GetBalanceAndLimitByUserId(db, userId)
		mu.Lock()
		defer mu.Unlock()
		if err != nil {
			resultErr = err
			return
		}
		balance = b
	}()

	wg.Wait()

	if resultErr != nil {
		return nil, resultErr
	}

	return &dto.ExtractOutputDTO{
		Balance:          *balance,
		LastTransactions: *transactions,
	}, nil
}
