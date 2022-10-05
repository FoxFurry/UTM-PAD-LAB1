package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/catalogue/internal/store"
	catalogue "pad/services/catalogue/proto"
)

type Catalogue struct {
	datastore store.Catalogue

	catalogue.UnimplementedCatalogueServer
}

func NewCatalogueServer(db store.Catalogue) catalogue.CatalogueServer {
	return &Catalogue{
		datastore: db,
	}
}

func (c *Catalogue) AddListing(ctx context.Context, l *catalogue.AddListingRequest) (*emptypb.Empty, error) {
	err := c.datastore.AddListing(ctx, &store.Listing{
		Title:        l.GetListing().Title,
		Description:  l.GetListing().Description,
		ThumbnailURL: l.GetListing().ThumbnailUrl,
		AuthorID:     69, // TODO(Arthur): Change to actual author id once auth header is added
	})
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *Catalogue) GetAllListings(ctx context.Context, _ *emptypb.Empty) (*catalogue.GetAllListingsResponse, error) {
	listings, err := c.datastore.GetAllListings(ctx)
	if err != nil {
		return nil, err
	}

	return &catalogue.GetAllListingsResponse{
		Listings: listingStoreToProto(listings),
	}, nil
}

func (c *Catalogue) GetListingByTitle(ctx context.Context, req *catalogue.GetListingByTitleRequest) (*catalogue.GetListingByTitleResponse, error) {
	listing, err := c.datastore.GetListingByTitle(ctx, req.GetTitle())
	if err != nil {
		return nil, err
	}

	return &catalogue.GetListingByTitleResponse{
		Listing: &catalogue.Listing{
			Title:        listing.Title,
			Description:  listing.Description,
			ThumbnailUrl: listing.ThumbnailURL,
		},
	}, nil
}

func listingStoreToProto(l []store.Listing) []*catalogue.Listing {
	var buffer []*catalogue.Listing

	for idx := range l {
		buffer = append(buffer, &catalogue.Listing{
			Title:        l[idx].Title,
			Description:  l[idx].Description,
			ThumbnailUrl: l[idx].ThumbnailURL,
		})
	}

	return buffer
}
