package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
