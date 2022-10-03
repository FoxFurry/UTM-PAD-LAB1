package store

import (
	"database/sql"
)

type Catalogue interface {
}

type catalogue struct {
	db *sql.DB
}

func NewCatalogueStore(db *sql.DB) Catalogue {
	return catalogue{
		db: db,
	}
}
