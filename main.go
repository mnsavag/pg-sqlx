package main

import (
	"context"
	"fmt"

	"github.com/mnsavag/pg-sqlx/repository"

	"github.com/google/uuid"
)

const idStr = "20823ae8-ecde-4634-9a2d-7be8f0deeaba"

func main() {
	conn, _ := repository.NewSqliteConn("./example.db")
	repo := repository.NewSqliteRepository(conn)

	//addeck(repo)
	//getDeckById(repo, idStr)
	//deleteDeck(repo, idStr)

	//update
	uid, _ := uuid.Parse(idStr)
	//updateDeckDataTopic(repo, uid)
	//updateDeckDataDesc(repo, uid)
	//updateDeckDataLinks(repo, uid)
	updateDeckDataAll(repo, uid)

}

func addeck(repo *repository.SqliteRepository) string {
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
		return ""
	}

	fmt.Printf("Add Deck Done %s\n", id.String())
	return id.String()
}

func getDeckById(repo *repository.SqliteRepository, id string) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	deck, err := repo.GetDeckById(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(deck)
}

func deleteDeck(repo *repository.SqliteRepository, id string) {
	uid, _ := uuid.Parse(id)
	ctx := context.Background()

	err := repo.DeleteDeck(ctx, uid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Delete deck with id: %s, check db\n", id)
}

func updateDeckDataTopic(repo *repository.SqliteRepository, id uuid.UUID) {
	topic := "NewField"

	err := repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Topic: &topic,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Topic with deck %s success\n", id.String())
}

func updateDeckDataDesc(repo *repository.SqliteRepository, id uuid.UUID) {
	desc := "NewField"

	err := repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Description: &desc,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Description with deck %s success\n", id.String())
}

func updateDeckDataLinks(repo *repository.SqliteRepository, id uuid.UUID) {
	links := []string{"NewLink1", "NewLink2"}

	err := repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
		Links: links,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Update Links with deck %s success\n", id.String())
}

func updateDeckDataAll(repo *repository.SqliteRepository, id uuid.UUID) {
	topic := "all"
	desc := "all"
	links := []string{"all1", "all2"}

	err := repo.UpdateDeck(context.Background(), id, repository.UpdateDeckData{
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
