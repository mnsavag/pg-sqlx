package main

import (
	"context"
	"fmt"

	"github.com/mnsavag/pg-sqlx/repository"

	"github.com/google/uuid"
)

func main() {
	conn, _ := repository.NewSqliteConn("./example.db")
	repo := repository.NewSqliteRepository(conn)

	// Add Deck
	addeck(repo)
}

func addeck(repo *repository.SqliteRepository) {
	// Add Deck
	ctx := context.Background()
	topic, desc := "AddDeck", "AddDeck"
	deck := repository.Deck{
		Id:          uuid.New(),
		Topic:       &topic,
		Description: &desc,
		Links:       []string{"link1", "link2"},
	}
	id, err := repo.AddDeck(ctx, &deck)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Add Deck Done %s\n", id.String())
}

//https://github.com/golang-migrate/migrate
