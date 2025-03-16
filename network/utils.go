package network

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	c.JSON(http.StatusOK, result)
}

func (n *Network) badRequestResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusBadRequest, result)
}

func (n *Network) unAuthorizedResponse(c *gin.Context, result interface{}) {
	c.JSON(http.StatusUnauthorized, result)
}

func (n *Network) forbiddenRequest(c *gin.Context, result interface{}) {
	c.JSON(http.StatusForbidden, result)
}
