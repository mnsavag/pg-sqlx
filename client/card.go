package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mnsavag/pg-sqlx/repository"
)

func (c *Client) AddCard(deckId uuid.UUID) string {
	// Add Card
	ctx := context.Background()
	question, answer := "AddCard", "AddCard"
	card := repository.Card{
		Id:       uuid.New(),
		Question: &question,
		Answer:   &answer,
	}

	id, err := c.repo.AddCard(ctx, deckId, &card)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Printf("Add Card Done %s\n", id.String())
	return id.String()
}

func (c *Client) DeleteCard(id string) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	err := c.repo.DeleteCard(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Delete card with id: %s, check db\n", id)
}

func (c *Client) UpdateCardDataAll(id uuid.UUID) {
	question := "all"
	answer := "all"

	err := c.repo.UpdateCard(context.Background(), id, repository.UpdateCardData{
		Question: &question,
		Answer:   &answer,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update all fields with card %s success\n", id.String())
}
