package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	catalogue "pad/services/catalogue/proto"
)

type Catalogue struct {
	catalogue.UnimplementedCatalogueServer
}

func NewCatalogueServer() catalogue.CatalogueServer {
	return &Catalogue{}
}

func (c *Catalogue) AddListing(context.Context, *catalogue.Listing) (*emptypb.Empty, error) {

}

func (c *Catalogue) GetAllListings(context.Context, *emptypb.Empty) (*catalogue.GetListingsResponse, error) {

}
