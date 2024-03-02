package database

import (
	"database/sql"

	"log"

	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateTransaction(*dto.TransactionInputDTO) error
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=rinha_de_backend_2024_q1 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("error on open connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("error on ping connection: %v", err)
		return nil, err
	}

	return &PostgresStorage{
		db: db,
	}, nil
}

func (p *PostgresStorage) CreateTransaction(transaction *dto.TransactionInputDTO) (*dto.TransactionDTO, error) {
	query := "INSERT INTO public.transacoes (valor, tipo, descricao, cliente_id) VALUES ($1, $2, $3, $4) RETURNING id, valor, tipo, descricao, cliente_id"

	stmt, err := p.db.Prepare(query)
	if err != nil {
		log.Printf("error on prepare statement: %v", err)
		return nil, err
	}

	var transactionCreated dto.TransactionDTO
	err = stmt.QueryRow(transaction.Valor, transaction.Tipo, transaction.Descricao).Scan(&transactionCreated.ID, &transactionCreated.Valor, &transactionCreated.Tipo, &transactionCreated.Descricao)
	if err != nil {
		log.Printf("error on query row: %v", err)
		return nil, err
	}

	return &transactionCreated, nil
}
