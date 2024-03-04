package database

import (
	"database/sql"
	"time"

	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kleytonsolinho/rinha-de-backend-2024-q1/internal/infra/dto"
)

type Storage interface {
	CreateTransaction(*dto.TransactionInputDTO) error
}

func NewMySQLStorage() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/mysql-db?parseTime=true")
	if err != nil {
		log.Printf("error on open connection: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Printf("error on ping connection: %v", err)
		return nil, err
	}

	return db, nil
}

func CreateTransaction(db *sql.DB, transaction *dto.TransactionInputDTO) error {
	query := "INSERT INTO transacoes (valor, tipo, descricao, cliente_id, realizado_em) VALUES (?, ?, ?, ?, ?)"
	queryUpdateBalance := "UPDATE clientes SET saldo = saldo + ? WHERE id = ?"

	if transaction.Tipo == "d" {
		transaction.Valor = transaction.Valor * -1
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("error on prepare statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.Valor, transaction.Tipo, transaction.Descricao, transaction.ClienteID, time.Now().Format("2006-01-02T15:04:05.999999Z"))
	if err != nil {
		log.Printf("error on query row: %v", err)
		return err
	}

	stmt, err = db.Prepare(queryUpdateBalance)
	if err != nil {
		log.Printf("error on prepare statement: %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(transaction.Valor, transaction.ClienteID)
	if err != nil {
		log.Printf("error on query row: %v", err)
		return err
	}

	return nil
}

func GetBalanceAndLimitByUserId(db *sql.DB, userId int64) (*dto.Balance, error) {
	query := "SELECT saldo, limite FROM clientes WHERE id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("error on prepare statement: %v", err)
		return nil, err
	}
	defer stmt.Close()

	var balance dto.Balance
	row := stmt.QueryRow(userId)
	err = row.Scan(&balance.Total, &balance.Limit)
	if err != nil {
		log.Printf("error on scan row: %v", err)
		return nil, err
	}

	balance.CreatedAt = time.Now().Format("2006-01-02T15:04:05.999999Z")

	return &balance, nil
}

func GetLastTransactionsByUserId(db *sql.DB, id int64) ([]dto.LastTransaction, error) {
	query := "SELECT valor, tipo, descricao, realizado_em FROM transacoes WHERE cliente_id = ? ORDER BY realizado_em DESC LIMIT 10"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Printf("error on prepare statement: %v", err)
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("error on query row: %v", err)
		return nil, err
	}
	defer rows.Close()

	var transactions []dto.LastTransaction
	for rows.Next() {
		var transaction dto.LastTransaction
		err = rows.Scan(&transaction.Valor, &transaction.Tipo, &transaction.Descricao, &transaction.RealizadoEm)
		if err != nil {
			log.Printf("error on scan row: %v", err)
			return nil, err
		}

		if transaction.Tipo == "d" {
			transaction.Valor = transaction.Valor * -1
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
