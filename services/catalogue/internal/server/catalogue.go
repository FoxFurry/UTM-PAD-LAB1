package server

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/cache/services/cache"
	"pad/services/catalogue/internal/store"
	"pad/services/catalogue/services/catalogue"
)

type Catalogue struct {
	datastore store.Catalogue
	cache     cache.CacheClient

	catalogue.UnimplementedCatalogueServer
}

func NewCatalogueServer(db store.Catalogue, cacheService cache.CacheClient) catalogue.CatalogueServer {
	return &Catalogue{
		datastore: db,
		cache:     cacheService,
	}
}

func (c *Catalogue) AddListing(ctx context.Context, req *catalogue.AddListingRequest) (*emptypb.Empty, error) {
	err := c.datastore.AddListing(ctx, &store.Listing{
		Title:        req.GetListing().Title,
		Description:  req.GetListing().Description,
		ThumbnailURL: req.GetListing().ThumbnailUrl,
		AuthorID:     69, // TODO(Arthur): Change to actual author id once auth header is added
	})
	if err != nil {
		return nil, err
	}

	if _, err := c.cache.AddListing(ctx, &cache.AddListingRequest{Listing: listingProtoToCache(req.GetListing())}); err != nil {
		log.Printf("failed to cache listing: %v\n", err)
		return nil, err
	}
	log.Printf("cached listing by title: %s\n", req.GetListing().GetTitle())

	return &emptypb.Empty{}, nil
}

func (c *Catalogue) GetAllListings(ctx context.Context, _ *emptypb.Empty) (*catalogue.GetAllListingsResponse, error) {
	listings, err := c.datastore.GetAllListings(ctx)
	if err != nil {
		return nil, err
	}

	return &catalogue.GetAllListingsResponse{
		Listings: listingsStoreToProto(listings),
	}, nil
}

func (c *Catalogue) GetListingByTitle(ctx context.Context, req *catalogue.GetListingByTitleRequest) (*catalogue.GetListingByTitleResponse, error) {
	if cachedListing, err := c.cache.GetListingByTitle(ctx, &cache.GetListingByTitleRequest{Title: req.Title}); err == nil {
		log.Printf("found cached listing for title: %s\n", req.Title)

		return &catalogue.GetListingByTitleResponse{
			Listing: listingCacheToProto(cachedListing.Listing),
		}, nil
	}

	log.Printf("cache for %s not found, looking in db\n", req.Title)

	listing, err := c.datastore.GetListingByTitle(ctx, req.GetTitle())
	if err != nil {
		return nil, err
	}

	catalogueListing := &catalogue.Listing{
		Title:        listing.Title,
		Description:  listing.Description,
		ThumbnailUrl: listing.ThumbnailURL,
	}

	if _, err := c.cache.AddListing(ctx, &cache.AddListingRequest{Listing: listingProtoToCache(catalogueListing)}); err != nil {
		log.Printf("failed to cache listing: %v\n", err)
		return nil, err
	}
	log.Printf("cached listing by title: %s\n", req.Title)

	return &catalogue.GetListingByTitleResponse{
		Listing: catalogueListing,
	}, nil
}

func listingsStoreToProto(l []store.Listing) []*catalogue.Listing {
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

func listingCacheToProto(l *cache.Listing) *catalogue.Listing {
	return &catalogue.Listing{
		Title:        l.Title,
		Description:  l.Description,
		ThumbnailUrl: l.ThumbnailUrl,
	}
}

func listingProtoToCache(l *catalogue.Listing) *cache.Listing {
	return &cache.Listing{
		Title:        l.Title,
		Description:  l.Description,
		ThumbnailUrl: l.ThumbnailUrl,
	}
}
