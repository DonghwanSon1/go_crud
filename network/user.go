package network

import (
	"github.com/gin-gonic/gin"
	"go_crud/service"
	"go_crud/types"
	"go_crud/types/errors"
	"go_crud/types/usersInfo"
	"sync"
)

var (
	userRouterInit     sync.Once
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
		}

		router.registerPOST("/signup", userRouterInstance.signup)
		router.registerPOST("/login", userRouterInstance.login)
		router.registerGET("/", userRouterInstance.get)
		router.registerUPDATE("/", userRouterInstance.update)
		router.registerDELETE("/", userRouterInstance.delete)
	})
	return userRouterInstance
}

func (u *userRouter) signup(c *gin.Context) {
	var req usersInfo.SignupRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Signup(req.CreateUsersInfo()); err != nil {
		u.router.failedResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("Singup 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &usersInfo.SignupUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) login(c *gin.Context) {
	var req usersInfo.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if result, err := u.userService.Login(req); err != nil {
		u.router.failedResponse(c, &errors.ErrorResponse{
			ApiResponse: types.NewApiResponse("Login 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &usersInfo.LoginResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
			Token:       result,
		})
	}
}

func (u *userRouter) get(c *gin.Context) {
	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	var req types.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.ErrorResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Update(req.Email, req.UpdateAge); err != nil {
		u.router.failedResponse(c, &types.ErrorResponse{
			ApiResponse: types.NewApiResponse("Update 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *userRouter) delete(c *gin.Context) {
	var req types.DeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("바인딩 오류 입니다.", -1, err.Error()),
		})
	} else if err = u.userService.Delete(req.ToUser()); err != nil {
		u.router.failedResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("Delete 에러 입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}
