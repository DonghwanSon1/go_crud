package network

import "github.com/gin-gonic/gin"

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

// Response 형태 맞추기 위한 유틸 함수 입니다.
func (n *Network) okResponse(c *gin.Context, result interface{}) {
	c.JSON(200, result)
}

func (n *Network) failedResponse(c *gin.Context, result interface{}) {
	c.JSON(400, result)
}
