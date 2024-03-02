package entity

type TransactionEntity struct {
	Valor     int64
	Tipo      string
	Descricao string
}

func NewTransaction(valor int64, tipo string, descricao string) *TransactionEntity {
	return &TransactionEntity{
		Valor:     valor,
		Tipo:      tipo,
		Descricao: descricao,
	}
}
