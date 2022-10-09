package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pad/services/catalogue/services/catalogue"
)

func NewCatalogueClient(address string) (catalogue.CatalogueClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return catalogue.NewCatalogueClient(conn), nil
}
