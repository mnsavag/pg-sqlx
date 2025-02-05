package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

const (
	decksTable     = "decks"
	cardsTable     = "cards"
	deckCardsTable = "deck_cards"
)

func NewSqliteConn(storagePath string) (*sqlx.DB, error) {
	conn, err := sqlx.Connect("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return conn, nil
}
