package store

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Catalogue interface {
	GetAllListings(context.Context) ([]Listing, error)
	AddListing(context.Context, *Listing) error
	GetListingByTitle(context.Context, string) (*Listing, error)
}

type catalogue struct {
	db *sqlx.DB
}

func NewCatalogueStore(db *sqlx.DB) Catalogue {
	return &catalogue{
		db: db,
	}
}
