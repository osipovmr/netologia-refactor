package db

import "database/sql"

// DB описывает контракт инициализации и миграции БД
type DB interface {
	Open(path string) (*sql.DB, error)
	Migrate(db *sql.DB) error
}
