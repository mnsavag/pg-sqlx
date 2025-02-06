package main

import (
	"github.com/mnsavag/pg-sqlx/client"
	"github.com/mnsavag/pg-sqlx/repository"

	"github.com/google/uuid"
)

func main() {
	conn, _ := repository.NewSqliteConn("./example.db")
	repo := repository.NewSqliteRepository(conn)
	client := client.NewClient(repo)

	deckOperations(client)
}

func deckOperations(c *client.Client) {
	idStr := "d545d2ef-c905-4eb0-a69c-44b75a4fd037"
	uid, _ := uuid.Parse(idStr)

	//c.AddDeck()
	//c.GetDeckById(idStr)
	//c.DeleteDeck(idStr)

	//update
	//c.UpdateDeckDataTopic(uid)
	//c.UpdateDeckDataDesc(uid)
	//c.UpdateDeckDataLinks(uid)
	c.UpdateDeckDataAll(uid)
}
