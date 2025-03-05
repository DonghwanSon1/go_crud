package network

import "github.com/gin-gonic/gin"

type Network struct {
	engin *gin.Engine
}

func NewNetwork() *Network {
	r := &Network{
		engin: gin.New(),
	}

	newUserRouter(r)
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engin.Run(port)
}

// register 유틸 함수들
func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engin.DELETE(path, handler...)
}
