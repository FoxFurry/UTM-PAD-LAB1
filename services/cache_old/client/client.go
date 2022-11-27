package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pad/services/cache_old/services/cache"
)

func NewCacheClient(address string) (cache.CacheClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return cache.NewCacheClient(conn), nil
}
