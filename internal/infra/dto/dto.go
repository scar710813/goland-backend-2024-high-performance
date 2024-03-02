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

type Balance struct {
	Total     int64  `json:"total"`
	CreatedAt string `json:"data_extrato"`
	Limit     int64  `json:"limite"`
}

type LastTransaction struct {
	Valor       int64  `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadoEm string `json:"realizado_em"`
}

type ExtractOutputDTO struct {
	Balance          Balance           `json:"saldo"`
	LastTransactions []LastTransaction `json:"ultimas_transacoes"`
}
