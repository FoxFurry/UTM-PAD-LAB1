package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"pad/services/cache/services/cache"
)

const (
	titlePrefix = "title:"
	idPrefix    = "id:"
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
	c.listings[byTitle(req.Listing.Title)] = req.GetListing()
	c.listings[byID(req.GetId())] = req.GetListing()

	return &emptypb.Empty{}, nil
}

func (c *Cache) GetListingByID(ctx context.Context, req *cache.GetListingByIDRequest) (*cache.GetListingByIDResponse, error) {
	listing, ok := c.listings[byID(req.GetId())]
	if !ok {
		return nil, status.Error(codes.NotFound, "listing not found")
	}

	return &cache.GetListingByIDResponse{
		Listing: listing,
	}, nil
}

func (c *Cache) GetListingByTitle(ctx context.Context, req *cache.GetListingByTitleRequest) (*cache.GetListingByTitleResponse, error) {
	listing, ok := c.listings[byTitle(req.GetTitle())]
	if !ok {
		return nil, status.Error(codes.NotFound, "listing not found")
	}

	return &cache.GetListingByTitleResponse{
		Listing: listing,
	}, nil
}

func byTitle(title string) string {
	return titlePrefix + title
}

func byID(id uint32) string {
	return fmt.Sprintf("%s%d", idPrefix, id)
}
