package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/catalogue/internal/store"
	catalogue "pad/services/catalogue/proto"
)

type Catalogue struct {
	db store.Catalogue

	catalogue.UnimplementedCatalogueServer
}

func NewCatalogueServer(db store.Catalogue) catalogue.CatalogueServer {
	return &Catalogue{
		db: db,
	}
}

func (c *Catalogue) AddListing(context.Context, *catalogue.Listing) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (c *Catalogue) GetAllListings(context.Context, *emptypb.Empty) (*catalogue.GetListingsResponse, error) {
	return &catalogue.GetListingsResponse{}, nil
}
