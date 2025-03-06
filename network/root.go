package network

import (
	"github.com/gin-gonic/gin"
	"go_crud/service"
)

type Network struct {
	engin *gin.Engine

	service *service.Service
}

func NewNetwork(service *service.Service) *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r, service.User)
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}
