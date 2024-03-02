package dto

type TransactionDTO struct {
	ID        int64  `json:"id"`
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
	ClienteID int64  `json:"cliente_id"`
}

type TransactionInputDTO struct {
	Valor     int64  `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
	ClienteID int64  `json:"cliente_id"`
}

type TransactionOutputDTO struct {
	Limite int64 `json:"limite"`
	Saldo  int64 `json:"saldo"`
}
