package repository

import (
	"database/sql"
	"netologia-refactor/internal/model"
)

// SQLiteOrderRepository реализует RepositoryWriter для SQLite
type SQLiteOrderRepository struct {
	db *sql.DB
}

func NewSQLiteOrderRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{db: db}
}

func (r *SQLiteOrderRepository) CreateOrder(order model.Order) error {
	_, err := r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, order.Products, order.Total, order.Status,
	)
	return err
}
