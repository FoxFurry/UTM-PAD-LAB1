package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/cache/services/cache"
	"pad/services/catalogue/internal/store"
	"pad/services/catalogue/services/catalogue"
	"pad/services/common/apiauth"
)

type Catalogue struct {
	datastore store.Catalogue
	cache     cache.CacheClient

	catalogue.UnimplementedCatalogueServer
}

var (
	catalogueKey = "Hd7h4kaz9A)j4hf6G@#ts78tynblx"
	cacheKey     = "&k4kaHAHJb3b98a8h6jk5dsfMK8hb2s"
)

func NewCatalogueServer(db store.Catalogue, cacheService cache.CacheClient) catalogue.CatalogueServer {
	return &Catalogue{
		datastore: db,
		cache:     cacheService,
	}
}

func (c *Catalogue) AddListing(ctx context.Context, req *catalogue.AddListingRequest) (*emptypb.Empty, error) {
	err := c.datastore.AddListing(ctx, &store.Listing{
		Title:        req.GetListing().GetTitle(),
		Description:  req.GetListing().GetDescription(),
		ThumbnailURL: req.GetListing().GetThumbnailUrl(),
		AuthorID:     69, // TODO(Arthur): Change to actual author id once auth header is added
	})
	if err != nil {
		return nil, err
	}

	ctx = apiauth.SetRequestMetadataKey(ctx, catalogueKey)

	var cacheMetadata metadata.MD
	if _, err := c.cache.AddListing(ctx, &cache.AddListingRequest{Listing: listingProtoToCache(req.GetListing())}, grpc.Header(&cacheMetadata)); err != nil {
		log.Printf("failed to cache listing: %v\n", err)
		return nil, err
	}

	if err := apiauth.ValidateMetadataKey(cacheMetadata, cacheKey); err != nil {
		return nil, fmt.Errorf("failed to authenticate cache: %w", err)
	}

	log.Printf("cached listing by title: %s\n", req.GetListing().GetTitle())
	log.Printf("added to db listing by title: %s\n", req.GetListing().GetTitle())

	return &emptypb.Empty{}, nil
}

func (c *Catalogue) GetAllListings(ctx context.Context, _ *emptypb.Empty) (*catalogue.GetAllListingsResponse, error) {
	listings, err := c.datastore.GetAllListings(ctx)
	if err != nil {
		return nil, err
	}

	log.Printf("returning %d listings\n", len(listings))

	return &catalogue.GetAllListingsResponse{
		Listings: listingsStoreToProto(listings),
	}, nil
}

func (c *Catalogue) GetListingByTitle(ctx context.Context, req *catalogue.GetListingByTitleRequest) (*catalogue.GetListingByTitleResponse, error) {
	ctx = apiauth.SetRequestMetadataKey(ctx, catalogueKey)

	var cacheMetadata metadata.MD
	if cachedListing, err := c.cache.GetListingByTitle(ctx, &cache.GetListingByTitleRequest{Title: req.GetTitle()}, grpc.Header(&cacheMetadata)); err == nil {
		if err := apiauth.ValidateMetadataKey(cacheMetadata, cacheKey); err != nil {
			return nil, fmt.Errorf("failed to authenticate cache: %w", err)
		}

		log.Printf("found cached listing for title: %s\n", req.GetTitle())

		return &catalogue.GetListingByTitleResponse{
			Listing: listingCacheToProto(cachedListing.Listing),
		}, nil
	}

	log.Printf("cache for title %s not found, looking in db\n", req.GetTitle())

	listing, err := c.datastore.GetListingByTitle(ctx, req.GetTitle())
	if err != nil {
		log.Printf("title %s not found in database\n", req.GetTitle())
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
	log.Printf("cached listing by title: %s\n", req.GetTitle())

	return &catalogue.GetListingByTitleResponse{
		Listing: catalogueListing,
	}, nil
}

func (c *Catalogue) GetListingByID(ctx context.Context, req *catalogue.GetListingByIDRequest) (*catalogue.GetListingByIDResponse, error) {
	ctx = apiauth.SetRequestMetadataKey(ctx, catalogueKey)

	var cacheMetadata metadata.MD
	if cachedListing, err := c.cache.GetListingByID(ctx, &cache.GetListingByIDRequest{Id: req.GetId()}, grpc.Header(&cacheMetadata)); err == nil {
		if err := apiauth.ValidateMetadataKey(cacheMetadata, cacheKey); err != nil {
			return nil, fmt.Errorf("failed to authenticate cache: %w", err)
		}

		log.Printf("found cached listing for id: %d\n", req.GetId())

		return &catalogue.GetListingByIDResponse{
			Listing: listingCacheToProto(cachedListing.Listing),
		}, nil
	}

	log.Printf("cache for id %d not found, looking in db\n", req.GetId())

	listing, err := c.datastore.GetListingByID(ctx, req.GetId())
	if err != nil {
		log.Printf("id %d not found in database\n", req.GetId())
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
	log.Printf("cached listing by id: %d\n", req.GetId())

	return &catalogue.GetListingByIDResponse{
		Listing: catalogueListing,
	}, nil
}

func (c *Catalogue) Heartbeat(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
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
