package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func (r *SqliteRepository) AddCard(ctx context.Context, deckId uuid.UUID, card *Card) (uuid.UUID, error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return uuid.Nil, err
	}

	insertCardsQuery := fmt.Sprintf("INSERT INTO %s (id, question, answer) VALUES (:id, :question, :answer)", cardsTable)
	insertCardsQuery, args, err := sqlx.Named(insertCardsQuery, map[string]interface{}{
		"id":       card.Id.String(),
		"question": card.Question,
		"answer":   card.Answer,
		"deck_id":  deckId,
	})
	if err != nil {
		tx.Rollback()
		return uuid.Nil, errors.Errorf("cant create card: %s", err.Error())
	}

	_, err = tx.ExecContext(ctx, insertCardsQuery, args...)
	if err != nil {
		return uuid.Nil, errors.Errorf("cant create card: %s", err.Error())
	}

	insertDeckCardsQuery := fmt.Sprintf("INSERT INTO %s (deck_id, card_id) VALUES (:deck_id, :card_id)", deckCardsTable)
	insertDeckCardsQuery, args, err = sqlx.Named(insertDeckCardsQuery, map[string]interface{}{
		"deck_id": deckId.String(),
		"card_id": card.Id.String(),
	})
	if err != nil {
		tx.Rollback()
		return uuid.Nil, errors.Errorf("cant create card: %s", err.Error())
	}

	_, err = tx.ExecContext(ctx, insertDeckCardsQuery, args...)
	if err != nil {
		tx.Rollback()
		return uuid.Nil, errors.Errorf("cant create card: %s", err.Error())
	}

	return card.Id, tx.Commit()
}

func (r *SqliteRepository) DeleteCard(ctx context.Context, id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}

	delCardQuery := fmt.Sprintf("DELETE FROM %s WHERE id=:1", cardsTable)
	_, err = tx.ExecContext(ctx, delCardQuery, id.String())
	if err != nil {
		tx.Rollback()
		return errors.Errorf("cant delete card: %s", err.Error())
	}

	delDeckCardsQuery := fmt.Sprintf("DELETE FROM %s WHERE card_id=:1", deckCardsTable)
	_, err = tx.ExecContext(ctx, delDeckCardsQuery, id.String())
	if err != nil {
		tx.Rollback()
		return errors.Errorf("cant delete card: %s", err.Error())
	}

	return tx.Commit()
}

func (r *SqliteRepository) UpdateCard(ctx context.Context, id uuid.UUID, fieldsUpdate UpdateCardData) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if fieldsUpdate.Question != nil {
		setValues = append(setValues, fmt.Sprintf("question=:%d", argId))
		args = append(args, &fieldsUpdate.Question)
		argId++
	}

	if fieldsUpdate.Answer != nil {
		setValues = append(setValues, fmt.Sprintf("answer=:%d", argId))
		args = append(args, &fieldsUpdate.Answer)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=:%d",
		cardsTable, setQuery, argId)
	args = append(args, id.String())

	_, err := r.db.ExecContext(ctx, query, args...)
	return err
}
