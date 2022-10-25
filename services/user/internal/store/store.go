package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateCredentials(context.Context, *Credentials) error
	GetCredentials(context.Context, string) (*Credentials, error)
}

type user struct {
	db *sqlx.DB
}

func NewUserStore(db *sqlx.DB) User {
	return &user{
		db: db,
	}
}
