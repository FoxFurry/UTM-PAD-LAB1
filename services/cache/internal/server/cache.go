package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/cache/services/cache"
	"pad/services/common/apiauth"
)

const (
	titlePrefix  = "title:"
	idPrefix     = "id:"
	catalogueKey = "Hd7h4kaz9A)j4hf6G@#ts78tynblx"
	cacheKey     = "&k4kaHAHJb3b98a8h6jk5dsfMK8hb2s"
)

type Cache struct {
	cache.UnimplementedCacheServer

	listings map[string]*cache.Listing
}

func NewCacheServer() cache.CacheServer {
	return &Cache{
		listings: make(map[string]*cache.Listing),
	}
}

func (c *Cache) AddListing(ctx context.Context, req *cache.AddListingRequest) (*emptypb.Empty, error) {
	if err := apiauth.ValidateContextKey(ctx, catalogueKey); err != nil {
		return nil, fmt.Errorf("failed to authenticate catalogue: %w", err)
	}

	c.listings[byTitle(req.Listing.Title)] = req.GetListing()

	log.Printf("succesfully cached listing with title %s\n", req.Listing.Title)

	if req.Id != 0 {
		c.listings[byID(req.GetId())] = req.GetListing()
		log.Printf("succesfully cached listing by id %d\n", req.Id)
	}

	ctx, err := apiauth.SetResponseMetadataKey(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (c *Cache) GetListingByID(ctx context.Context, req *cache.GetListingByIDRequest) (*cache.GetListingByIDResponse, error) {
	if err := apiauth.ValidateContextKey(ctx, catalogueKey); err != nil {
		return nil, fmt.Errorf("failed to authenticate catalogue: %w", err)
	}

	listing, ok := c.listings[byID(req.GetId())]
	if !ok {
		log.Printf("listing with id %d not found\n", req.Id)
		return nil, status.Error(codes.NotFound, "listing not found")
	}

	log.Printf("listing with id %d was found\n", req.Id)

	ctx, err := apiauth.SetResponseMetadataKey(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	return &cache.GetListingByIDResponse{
		Listing: listing,
	}, nil
}

func (c *Cache) GetListingByTitle(ctx context.Context, req *cache.GetListingByTitleRequest) (*cache.GetListingByTitleResponse, error) {
	if err := apiauth.ValidateContextKey(ctx, catalogueKey); err != nil {
		return nil, fmt.Errorf("failed to authenticate catalogue: %w", err)
	}

	listing, ok := c.listings[byTitle(req.GetTitle())]
	if !ok {
		log.Printf("listing with title %s not found\n", req.Title)
		return nil, status.Error(codes.NotFound, "listing not found")
	}

	log.Printf("listing with title %s was found\n", req.Title)

	ctx, err := apiauth.SetResponseMetadataKey(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	return &cache.GetListingByTitleResponse{
		Listing: listing,
	}, nil
}

func (c *Cache) Heartbeat(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func byTitle(title string) string {
	return titlePrefix + title
}

func byID(id uint32) string {
	return fmt.Sprintf("%s%d", idPrefix, id)
}
