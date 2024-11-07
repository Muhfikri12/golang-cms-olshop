package repository

import (
	"database/sql"

	"github.com/Muhfikri12/golang-cms-olshop/model"
)

type RepoTransactionDB struct {
	DB *sql.DB
}

func NewRepoTransaction(db *sql.DB) RepoTransactionDB {
	return RepoTransactionDB{DB: db}
}

func (rt *RepoTransactionDB) Transaction() (*[]model.Transaction, error) {
	query := `
	SELECT c.name , 
       t.id , 
       t.status, 
       DATE(t.transaction_date)
	FROM transactions t
	JOIN users u ON t.user_id = u.id
	JOIN customers c ON u.id = c.user_id`

	rows, err := rt.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trxs := []model.Transaction{}
	for rows.Next() {
		trx := model.Transaction{} 
		if err := rows.Scan(&trx.Name, &trx.ID, &trx.Status, &trx.TransactionDate); err != nil {
			return nil, err
		}
		trxs = append(trxs, trx)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &trxs, nil
}