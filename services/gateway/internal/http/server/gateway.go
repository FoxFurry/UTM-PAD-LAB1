package server

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	catalogueClient "pad/services/catalogue/client"
	"pad/services/catalogue/services/catalogue"
	"pad/services/gateway/internal/http/httperr"
	"pad/services/gateway/internal/http/httpresponse"
	"pad/services/gateway/internal/http/middlewares"
	"pad/services/gateway/internal/models"
	"pad/services/locator/go-client/locator"
	"pad/services/user/services/user"
)

const (
	failsToCircuitBreak = 5
)

var (
	errGetParamMissing = fmt.Errorf("missing title or id parameter")
)

type Gateway struct {
	catalogue catalogue.CatalogueClient
	locator   locator.LocatorClient
	user      user.UserClient

	circuitStatus      bool
	circuitStatusMutex *sync.RWMutex
}

func NewGatewayServer(userClient user.UserClient, locatorClient locator.LocatorClient) *Gateway {
	return &Gateway{
		user:               userClient,
		locator:            locatorClient,
		circuitStatusMutex: &sync.RWMutex{},
		circuitStatus:      true,
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
	if !g.isCircuitOpen() {
		httpresponse.RespondInternalError(ctx, "circuit locked")
	}

	g.RelocateCatalogue(ctx)

	var listing models.Listing

	if err := ctx.BindJSON(&listing); err != nil {
		log.Printf("received bad request: %v\n", err)
		httperr.Handle(ctx, httperr.NewErrorBadRequest(err))
		return
	}

	for currentReroutes := 0; currentReroutes < failsToCircuitBreak; currentReroutes++ {
		_, err := g.catalogue.AddListing(ctx, &catalogue.AddListingRequest{
			Listing: &catalogue.Listing{
				Title:        listing.Title,
				Description:  listing.Description,
				ThumbnailUrl: listing.ThumbnailURL,
			},
		})
		if err == nil {
			httpresponse.RespondOK(ctx, gin.H{
				"successful": "listing created",
			})
			return
		}

		log.Printf("[ATTEMPT %d] could not create listing: %v\n", currentReroutes+1, err)
		log.Printf("Attempting reroute...\n")
		g.RelocateCatalogue(ctx)
	}

	log.Printf("ass explosion. CIRCUIT BREAKER WILL BE TRIGGERED HERE\n")
	g.closeCircuit()

	httperr.Handle(ctx, httperr.NewErrorInternal())
}

func (g *Gateway) GetListing(ctx *gin.Context) {
	if !g.isCircuitOpen() {
		httpresponse.RespondInternalError(ctx, "circuit locked")
		return
	}

	g.RelocateCatalogue(ctx)

	if id := ctx.Query("id"); id != "" {

		uid, _ := strconv.ParseUint(id, 10, 32)

		for currentReroutes := 0; currentReroutes < failsToCircuitBreak; currentReroutes++ {
			listing, err := g.catalogue.GetListingByID(ctx, &catalogue.GetListingByIDRequest{
				Id: uint32(uid),
			})
			if err == nil {
				httpresponse.RespondOK(ctx, listing.Listing)

				log.Printf("listing queried by id\n")

				return
			}

			log.Printf("[ATTEMPT %d] could not get listing by id: %v\n", currentReroutes, err)
			log.Printf("Attempting reroute...\n")

			g.RelocateCatalogue(ctx)
		}

		log.Printf("Bruh... Closing circuit\n")
		g.closeCircuit()
		httperr.Handle(ctx, httperr.NewErrorInternal())
		return
	}

	if title := ctx.Query("title"); title != "" {

		for currentReroutes := 0; currentReroutes < failsToCircuitBreak; currentReroutes++ {
			listing, err := g.catalogue.GetListingByTitle(ctx, &catalogue.GetListingByTitleRequest{
				Title: title,
			})
			if err == nil {
				httpresponse.RespondOK(ctx, listing.Listing)

				log.Printf("listing queried by title\n")

				return
			}

			log.Printf("[ATTEMPT %d] could not get listing by title: %v\n", currentReroutes, err)
			log.Printf("Attempting reroute...\n")

			g.RelocateCatalogue(ctx)
		}

		log.Printf("Bruh... Closing circuit\n")
		g.closeCircuit()
		httperr.Handle(ctx, httperr.NewErrorInternal())
		return
	}

	log.Printf("GetListing called but no parameters provided\n")

	httperr.Handle(ctx, httperr.NewErrorBadRequest(errGetParamMissing))
}

func (g *Gateway) RelocateCatalogue(ctx context.Context) {
	log.Printf("Asking locator service for catalogue address")

	catalogueAddress, err := g.locator.GetService(ctx, &locator.GetServiceRequest{
		Type: 1,
	})
	if err != nil {
		log.Fatal("failed to get catalogue address")
	}

	g.catalogue, err = catalogueClient.NewCatalogueClient(catalogueAddress.Address)
	if err != nil {
		log.Fatal("failed to connect to new catalogue server")
	}

	log.Printf("Locator successfully returned %s\n", catalogueAddress.Address)
}

func (g *Gateway) isCircuitOpen() bool {
	g.circuitStatusMutex.RLock()
	defer g.circuitStatusMutex.RUnlock()
	return g.circuitStatus
}

func (g *Gateway) closeCircuit() {
	g.closeCircuitStatus()

	go func() {
		log.Printf("Circuit cooldown started")

		timer := time.NewTimer(4 * time.Second)
		<-timer.C

		g.openCircuitStatus()

		log.Printf("Circuit cooldown finished")
	}()
}

func (g *Gateway) flipCircuitStatus() {
	g.circuitStatusMutex.Lock()
	defer g.circuitStatusMutex.Unlock()

	log.Printf("Circuit flipped")
	g.circuitStatus = !g.circuitStatus
}

func (g *Gateway) closeCircuitStatus() {
	g.circuitStatusMutex.Lock()
	defer g.circuitStatusMutex.Unlock()

	log.Printf("Circuit clossed")
	g.circuitStatus = false
}

func (g *Gateway) openCircuitStatus() {
	g.circuitStatusMutex.Lock()
	defer g.circuitStatusMutex.Unlock()

	log.Printf("Circuit openned")
	g.circuitStatus = true
}
