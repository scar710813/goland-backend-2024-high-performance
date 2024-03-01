package validators

import (
	"errors"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/webserver/handlers/transaction"
)

func NewTransactionValidator(value int64, tipo string, description string) (*transaction.TransactionInputDTO, error) {
	if value < 0 {
		return nil, errors.New("valor da transação não pode ser negativo")
	}
	if tipo != "c" && tipo != "d" {
		return nil, errors.New("tipo da transação inválido")
	}
	if len(description) < 1 || len(description) > 10 {
		return nil, errors.New("descrição da transação deve ter entre 1 e 10 caracteres")
	}

	return &transaction.TransactionValidator{
		value:       value,
		tipo:        tipo,
		description: description,
	}, nil
}
