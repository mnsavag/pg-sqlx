package main

import (
	"github.com/google/uuid"
	"github.com/mnsavag/pg-sqlx/client"
	"github.com/mnsavag/pg-sqlx/repository"
)

func main() {
	conn, _ := repository.NewSqliteConn("./example.db")
	repo := repository.NewSqliteRepository(conn)
	client := client.NewClient(repo)

	//deckOperations(client)
	cardOperations(client)
}

func deckOperations(c *client.Client) {
	// idStr := "d545d2ef-c905-4eb0-a69c-44b75a4fd037"
	// uid, _ := uuid.Parse(idStr)

	c.AddDeck()
	//c.GetDeckById(idStr)
	//c.DeleteDeck(idStr)

	//update
	//c.UpdateDeckDataTopic(uid)
	//c.UpdateDeckDataDesc(uid)
	//c.UpdateDeckDataLinks(uid)
	//c.UpdateDeckDataAll(uid)
}

func cardOperations(c *client.Client) {
	// deckId := "cd4f4bbf-6f1f-42a4-bb38-fa80e180c2b1"
	// deckUUID, _ := uuid.Parse(deckId)

	cardStrId := "b3e2ce7a-72c2-4112-a381-2351b0fbabad"
	cardUUID, _ := uuid.Parse(cardStrId)

	//c.AddCard(deckUUID)
	//c.DeleteCard(cardStrId)
	//c.UpdateCardDataAll(cardUUID)

	//update
	c.UpdateCardDataAll(cardUUID)
}
