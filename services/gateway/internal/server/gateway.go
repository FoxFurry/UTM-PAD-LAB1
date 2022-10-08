package server

type Gateway struct{}

func NewGatewayServer() *Gateway {
	return &Gateway{}
}

func (g *Gateway) Start(address string) error {

}
