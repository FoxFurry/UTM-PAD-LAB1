package server

import (
	"github.com/gin-gonic/gin"
	"pad/services/gateway/internal/models"
)

type Gateway struct{}

func NewGatewayServer() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Start(address string) error {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/listing", g.CreateListing)
		v1.GET("/listing", g.GetListing)
	}
}

func (g *Gateway) CreateListing(ctx *gin.Context) {
	var listing models.Listing

	if err := ctx.BindJSON(&listing); err != nil {

	}

}

func (g *Gateway) GetListing(ctx *gin.Context) {

}

func (g *Gateway) GetListingByTitle(ctx *gin.Context) {

}
