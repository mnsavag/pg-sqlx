package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteRepository struct {
	db *sqlx.DB
}

func NewSqliteRepository(db *sqlx.DB) *SqliteRepository {
	return &SqliteRepository{
		db: db,
	}
}
