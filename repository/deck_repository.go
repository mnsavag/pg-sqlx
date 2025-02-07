package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *SqliteRepository) AddDeck(ctx context.Context, deck *Deck) (uuid.UUID, error) {
	links, err := json.Marshal(deck.Links)
	if err != nil {
		return uuid.Nil, errors.Errorf("cant create deck: %s", err.Error())
	}

	query := fmt.Sprintf("INSERT INTO %s (id, topic, description, links) VALUES (:id, :topic, :description, :links)", decksTable)
	query, args, err := sqlx.Named(query, map[string]interface{}{
		"id":          deck.Id.String(),
		"topic":       &deck.Topic,
		"description": &deck.Description,
		"links":       links,
	})
	if err != nil {
		return uuid.Nil, errors.Errorf("cant create deck: %s", err.Error())
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return uuid.Nil, errors.Errorf("cant create deck: %s", err.Error())
	}

	return deck.Id, err
}

func (r *SqliteRepository) DeleteDeck(ctx context.Context, id uuid.UUID) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=:1", decksTable)
	_, err := r.db.ExecContext(ctx, query, id.String())

	return err
}

func (r *SqliteRepository) GetDeckById(ctx context.Context, id uuid.UUID) (Deck, error) {
	type DeckTable struct {
		Id          uuid.UUID
		Topic       *string
		Description *string
		Links       []byte
	}
	var deckTable DeckTable

	query := fmt.Sprintf(`SELECT dt.id, dt.topic, dt.description, dt.links FROM %s dt WHERE dt.id = :1`, decksTable)
	err := r.db.Get(&deckTable, query, id.String())
	if err != nil {
		return Deck{}, errors.Errorf("cant get deck: %s", err.Error())
	}

	var links []string
	err = json.Unmarshal(deckTable.Links, &links)
	if err != nil {
		return Deck{}, errors.Errorf("cant get deck: %s", err.Error())
	}

	return Deck{
		Id:          deckTable.Id,
		Topic:       deckTable.Topic,
		Description: deckTable.Description,
		Links:       links,
	}, err
}

func (r *SqliteRepository) UpdateDeck(ctx context.Context, id uuid.UUID, fieldsUpdate UpdateDeckData) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if fieldsUpdate.Topic != nil {
		setValues = append(setValues, fmt.Sprintf("topic=:%d", argId))
		args = append(args, &fieldsUpdate.Topic)
		argId++
	}

	if fieldsUpdate.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=:%d", argId))
		args = append(args, &fieldsUpdate.Description)
		argId++
	}

	if fieldsUpdate.Links != nil {
		links, err := json.Marshal(fieldsUpdate.Links)
		if err != nil {
			return errors.Errorf("cant create deck: %s", err.Error())
		}

		setValues = append(setValues, fmt.Sprintf("links=:%d", argId))
		args = append(args, links)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=:%d",
		decksTable, setQuery, argId)
	args = append(args, id.String())

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}
