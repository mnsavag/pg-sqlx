package client

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/mnsavag/pg-sqlx/repository"
)

func (c *Client) AddDeck() string {
	// Add Deck
	ctx := context.Background()
	topic, desc := "AddDeck", "AddDeck"
	deck := repository.Deck{
		Id:          uuid.New(),
		Topic:       &topic,
		Description: &desc,
		Links:       []string{"link1", "link2"},
	}
	id, err := c.repo.AddDeck(ctx, &deck)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Printf("Add Deck Done %s\n", id.String())
	return id.String()
}

func (c *Client) GetDeckById(id string) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	deck, err := c.repo.GetDeckById(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deck)
}

func (c *Client) DeleteDeck(id string) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	err := c.repo.DeleteDeck(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Delete deck with id: %s, check db\n", id)
}

func (c *Client) UpdateDeckDataTopic(id uuid.UUID) {
	topic := "NewField"

	err := c.repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Topic: &topic,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Topic with deck %s success\n", id.String())
}

func (c *Client) UpdateDeckDataDesc(id uuid.UUID) {
	desc := "NewField"

	err := c.repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Description: &desc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Description with deck %s success\n", id.String())
}

func (c *Client) UpdateDeckDataLinks(id uuid.UUID) {
	links := []string{"NewLink1", "NewLink2"}

	err := c.repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Links: links,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Links with deck %s success\n", id.String())
}

func (c *Client) UpdateDeckDataAll(id uuid.UUID) {
	topic := "all"
	desc := "all"
	links := []string{"all1", "all2"}

	err := c.repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Topic:       &topic,
		Description: &desc,
		Links:       links,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update all fields with deck %s success\n", id.String())
}
