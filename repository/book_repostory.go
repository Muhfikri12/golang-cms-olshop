package repository

import (
	"database/sql"
	"errors"

	"github.com/Muhfikri12/golang-cms-olshop/model"
)

type BookRepoDB struct {
	DB *sql.DB
}

func NewBookRepo(db *sql.DB) BookRepoDB {
	return BookRepoDB{DB: db}
}

func (r *BookRepoDB) CreateBookRepo(book *model.Books) error {
	query := `INSERT INTO items (name, category, author, price, discount, cover, pdf, stock) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	result, err := r.DB.Exec(query, book.Name, book.Category, book.Author, book.Price, book.Discount, book.Cover, book.Pdf, book.Stock)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return errors.New("failed to create book")
	}

	return nil
}

func (r *BookRepoDB) BookList() (*[]model.Books, error) {
	query := `SELECT id, name, category, author, price, discount FROM items`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []model.Books{}
	for rows.Next() {
		item := model.Books{}
		if err := rows.Scan(&item.ID, &item.Name, &item.Category, &item.Author, &item.Price, &item.Discount); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &items, nil
}
