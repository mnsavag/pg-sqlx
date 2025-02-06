package client

import "github.com/mnsavag/pg-sqlx/repository"

type Client struct {
	repo *repository.SqliteRepository
}

func NewClient(repo *repository.SqliteRepository) *Client {
	return &Client{
		repo: repo,
	}
}
