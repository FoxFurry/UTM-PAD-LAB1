package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"pad/services/catalogue/services/catalogue"
	"pad/services/gateway/internal/http/httperr"
	"pad/services/gateway/internal/http/httpresponse"
	"pad/services/gateway/internal/http/middlewares"
	"pad/services/gateway/internal/models"
	"pad/services/user/services/user"
)

var (
	errGetParamMissing = fmt.Errorf("missing title or id parameter")
)

type Gateway struct {
	catalogue catalogue.CatalogueClient
	user      user.UserClient
}

func NewGatewayServer(catalogueClient catalogue.CatalogueClient, userClient user.UserClient) *Gateway {
	return &Gateway{
		catalogue: catalogueClient,
		user:      userClient,
	}
}

func (g *Gateway) Start(address string) error {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/register", g.Register)
		v1.POST("/login", g.Login)
		v1.POST("/listing", middlewares.TokenAuthMiddleware(), g.CreateListing)
		v1.GET("/listing", middlewares.TokenAuthMiddleware(), g.GetListing)
	}

	if err := router.Run(address); err != nil {
		return err
	}

	return nil
}

func (g *Gateway) Register(ctx *gin.Context) {
	var credentials models.Credentials

	if err := ctx.BindJSON(&credentials); err != nil {
		log.Printf("received bad request: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorBadRequest(err))
		return
	}

	if _, err := g.user.Register(ctx, &user.RegisterRequest{
		Name:     credentials.Name,
		Password: credentials.Password,
	}); err != nil {
		log.Printf("failed to register: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorUnauthorized())
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"successful": "user registered",
	})
}

func (g *Gateway) Login(ctx *gin.Context) {
	var credentials models.Credentials

	if err := ctx.BindJSON(&credentials); err != nil {
		log.Printf("received bad request: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorBadRequest(err))
		return
	}

	token, err := g.user.Login(ctx, &user.LoginRequest{
		Name:     credentials.Name,
		Password: credentials.Password,
	})
	if err != nil {
		log.Printf("failed to register: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorUnauthorized())
		return
	}

	httpresponse.RespondOK(ctx, gin.H{
		"token": token.Token,
	})
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
