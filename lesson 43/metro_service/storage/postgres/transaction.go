package postgres

import (
	"atto/models"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (t *TransactionRepo) Create(transaction *models.CreateTransaction) error {
	_, err := t.Db.Exec("INSERT INTO transaction (id, card_id, amount, terminal_id, transaction_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		uuid.NewString(), transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType, time.Now(), time.Now())
	return err
}

func (t *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	var transaction models.Transaction
	err := t.Db.QueryRow("SELECT id, card_id, amount, terminal_id, transaction_type FROM transaction WHERE id = $1", id).
		Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType)
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (t *TransactionRepo) Update(transaction *models.Transaction) error {
	_, err := t.Db.Exec("UPDATE transaction SET card_id = $2, amount = $3, terminal_id = $4, transaction_type = $5 WHERE id = $1",
		transaction.Id, transaction.CardId, transaction.Amount, transaction.TerminalId, transaction.TransactionType)
	return err
}

func (t *TransactionRepo) Delete(id string) error {
	_, err := t.Db.Exec("UPDATE transaction SET deleted_at = $2 WHERE id = $1", id, time.Now())
	return err
}

func (t *TransactionRepo) GetAll() ([]models.Transaction, error) {
	rows, err := t.Db.Query("SELECT id, card_id, amount, terminal_id, transaction_type FROM transaction")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var transaction models.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.CardId, &transaction.Amount, &transaction.TerminalId, &transaction.TransactionType); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
