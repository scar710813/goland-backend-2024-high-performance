package validators

import (
	"errors"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

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
