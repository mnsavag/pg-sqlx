package repository

import "github.com/google/uuid"

type Deck struct {
	Id          uuid.UUID
	Topic       *string
	Description *string
	Links       []string
}

type UpdateDeckData struct {
	Topic       *string
	Description *string
	Links       []string
}

type UpdateCardData struct {
	Question *string
	Answer   *string
}

type Card struct {
	Id       uuid.UUID
	Question *string
	Answer   *string
}
