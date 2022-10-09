package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"pad/services/catalogue/services/catalogue"
	"pad/services/gateway/internal/http/httperr"
	"pad/services/gateway/internal/http/httpresponse"
	"pad/services/gateway/internal/models"
)

var (
	errGetParamMissing = fmt.Errorf("missing title or id parameter")
)

type Gateway struct {
	catalogue catalogue.CatalogueClient
}

func NewGatewayServer(catalogueClient catalogue.CatalogueClient) *Gateway {
	return &Gateway{
		catalogue: catalogueClient,
	}
}

func (g *Gateway) Start(address string) error {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/listing", g.CreateListing)
		v1.GET("/listing", g.GetListing)
	}

	if err := router.Run(address); err != nil {
		return err
	}

	return nil
}

func (g *Gateway) CreateListing(ctx *gin.Context) {
	var listing models.Listing

	if err := ctx.BindJSON(&listing); err != nil {
		log.Printf("received bad request: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorBadRequest(err))
		return
	}

	if _, err := g.catalogue.AddListing(ctx, &catalogue.AddListingRequest{
		Listing: &catalogue.Listing{
			Title:        listing.Title,
			Description:  listing.Description,
			ThumbnailUrl: listing.ThumbnailURL,
		},
	}); err != nil {
		log.Printf("could not create listing: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorInternal())
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"successful": "listing created",
	})
}

func (g *Gateway) GetListing(ctx *gin.Context) {
	if id := ctx.Query("id"); id != "" {
		log.Printf("listing queried by id\n")

		uid, _ := strconv.ParseUint(id, 10, 32)
		listing, err := g.catalogue.GetListingByID(ctx, &catalogue.GetListingByIDRequest{
			Id: uint32(uid),
		})
		if err != nil {
			log.Printf("could not get listing by id: %v\n", err)

			httperr.Handle(ctx, httperr.NewErrorNotFound())
			return
		}

		httpresponse.RespondOK(ctx, listing.Listing)
		return
	}

	if title := ctx.Query("title"); title != "" {
		log.Printf("listing queried by title\n")

		listing, err := g.catalogue.GetListingByTitle(ctx, &catalogue.GetListingByTitleRequest{
			Title: title,
		})
		if err != nil {
			log.Printf("could not get listing by title: %v\n", err)

			httperr.Handle(ctx, httperr.NewErrorNotFound())
			return
		}

		httpresponse.RespondOK(ctx, listing.Listing)
		return
	}

	log.Printf("GetListing called but no parameters provided\n")

	httperr.Handle(ctx, httperr.NewErrorBadRequest(errGetParamMissing))
}
