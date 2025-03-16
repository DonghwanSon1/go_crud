package network

import (
	"github.com/gin-gonic/gin"
	"go_crud/service"
	"go_crud/types"
	"go_crud/types/errors"
	"sync"
)

var (
	middlewareRouterInit     sync.Once
	middlewareRouterInstance *middlewareRouter
)

type middlewareRouter struct {
	router *Network
	// service
	middlewareService *service.Middleware
}

func (m *middlewareRouter) RegisterMiddleware() {
	m.router.engin.Use(m.tokenValidate)
}

func newMiddlewareRouter(router *Network, middlewareService *service.Middleware) *middlewareRouter {
	middlewareRouterInit.Do(func() {
		middlewareRouterInstance = &middlewareRouter{
			router:            router,
			middlewareService: middlewareService,
		}

		//router.registerGET("/middleware/token-validate", middlewareRouterInstance.tokenValidate)
	})
	return middlewareRouterInstance
}

func (m *middlewareRouter) tokenValidate(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token == "" {
		m.router.badRequestResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("토큰이 없습니다.", -1, "Authorization header is required."),
		})
		c.Abort()
		return
	}

	// Bearer 토큰이 있을 경우 "Bearer " 제거
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	if userId, err := m.middlewareService.ValidateToken(token); err != nil {
		m.router.unAuthorizedResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("유효하지 않은 토큰입니다. 재로그인 부탁드립니다.", -1, err.Error()),
		})
		c.Abort()
		return
	} else {
		c.Set("userId", userId)
		c.Next()
	}
}
